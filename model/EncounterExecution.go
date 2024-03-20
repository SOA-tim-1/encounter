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
	Entity
	EncounterId            int64                    `json:"encounterId" gorm:"column:EncounterID"`
	TouristId              int64                    `json:"touristId" gorm:"column:TouristID"`
	Status                 EncounterExecutionStatus `json:"status"`
	LastActivity           time.Time                `json:"lastActivity" gorm:"column:LastActivity"`
	LocationEntryTimestamp *time.Time               `json:"locationEntryTimestamp" gorm:"column:LocationEntryTimestamp"`
	LastPosition           Coordinate               `json:"lastPosition" gorm:"type:json;column:LastPosition"`
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
