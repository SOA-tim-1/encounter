package model

import (
	"errors"
	"time"
)

type EncounterExecutionStatus int32

const (
	ExecutionActive EncounterExecutionStatus = iota
	ExecutionCompleted
	ExecutionAbandoned
)

type EncounterExecution struct {
	ID                     int64                    `json:"id" gorm:"primaryKey:autoIncrement"`
	EncounterId            int64                    `json:"encounterId"`
	TouristId              int64                    `json:"touristId"`
	Status                 EncounterExecutionStatus `json:"status"`
	LastActivity           time.Time                `json:"lastActivity"`
	LocationEntryTimestamp *time.Time               `json:"locationEntryTimestamp"`
	LastPosition           Coordinate               `json:"lastPosition" gorm:"type:json"`
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
	ee.updateLastActivityInformation(nil)
	return nil
}

func (ee *EncounterExecution) CheckIfCompletedHiddenLocation(encounter Encounter, currentPosition Coordinate) error {
	if err := ee.validateIsActive(); err != nil {
		return err
	}
	ee.updateLastActivityInformation(&currentPosition)

	if encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp == nil {
		now := time.Now()
		ee.LocationEntryTimestamp = &now
	} else if !encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp != nil {
		ee.LocationEntryTimestamp = nil
	} else if encounter.IsWithinHiddenLocationRange(currentPosition) && ee.LocationEntryTimestamp != nil && ee.hasCompletedLocationEntryDelay() {
		ee.Complete(currentPosition)
	}
	return nil
}

func (ee *EncounterExecution) hasCompletedLocationEntryDelay() bool {
	return time.Since(*ee.LocationEntryTimestamp) >= time.Second*30
}

func (ee *EncounterExecution) Complete(currentPosition Coordinate) error {
	if err := ee.validateIsActive(); err != nil {
		return err
	}
	ee.Status = ExecutionCompleted
	ee.updateLastActivityInformation(&currentPosition)
	return nil
}

func (ee *EncounterExecution) updateLastActivityInformation(currentPosition *Coordinate) {
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
