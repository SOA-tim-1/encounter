package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncounterExecutionStatus int32

const (
	ExecutionActive EncounterExecutionStatus = iota
	ExecutionCompleted
	ExecutionAbandoned
)

type EncounterExecution struct {
	Entity
	EncounterId            primitive.ObjectID       `bson:"encounterId,omitempty" json:"encounterId"`
	TouristId              int64                    `bson:"touristId,omitempty" json:"touristId"`
	Status                 EncounterExecutionStatus `bson:"status,omitempty" json:"status"`
	LastActivity           time.Time                `bson:"lastActivity,omitempty" json:"lastActivity"`
	LocationEntryTimestamp *time.Time               `bson:"locationEntryTimestamp,omitempty" json:"locationEntryTimestamp"`
	LastPosition           Coordinate               `bson:"lastPosition,omitempty" json:"lastPosition"`
}

func NewEncounterExecution(encounter Encounter, touristId int64, status EncounterExecutionStatus, currentPosition Coordinate) (*EncounterExecution, error) {
	if !encounter.IsWithinRange(currentPosition) {
		return nil, errors.New("Not in encounter range")
	}

	return &EncounterExecution{
		EncounterId:  encounter.ID,
		TouristId:    touristId,
		Status:       status,
		LastActivity: time.Now(),
		LastPosition: currentPosition,
	}, nil
}

func (ee *EncounterExecution) Abandon() error {
	if err := ee.validateIsActive(); err != nil {
		return err
	}
	ee.Status = ExecutionAbandoned
	ee.UpdateLastActivityInformation(nil)
	return nil
}

func (ee *EncounterExecution) CheckIfCompletedHiddenLocation(encounter Encounter, currentPosition Coordinate) error {
	if err := ee.validateIsActive(); err != nil {
		return err
	}
	ee.UpdateLastActivityInformation(&currentPosition)

	if encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp == nil {
		now := time.Now()
		ee.LocationEntryTimestamp = &now
	} else if !encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp != nil {
		ee.LocationEntryTimestamp = nil
	} else if encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp != nil && ee.hasCompletedLocationEntryDelay() {
		ee.Complete(&currentPosition)
	}
	return nil
}

func (ee *EncounterExecution) hasCompletedLocationEntryDelay() bool {
	return time.Since(*ee.LocationEntryTimestamp) >= time.Second*30
}

func (ee *EncounterExecution) Complete(currentPosition *Coordinate) error {
	if err := ee.validateIsActive(); err != nil {
		return err
	}
	ee.Status = ExecutionCompleted
	ee.UpdateLastActivityInformation(currentPosition)
	return nil
}

func (ee *EncounterExecution) UpdateLastActivityInformation(currentPosition *Coordinate) {
	ee.LastActivity = time.Now()
	if currentPosition != nil {
		ee.LastPosition = *currentPosition
	}
}

func (ee *EncounterExecution) validateIsActive() error {
	if ee.Status != ExecutionActive {
		return errors.New("Encounter execution is not active")
	}
	return nil
}
