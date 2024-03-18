package dtos

import "time"

type EncounterExecutionStatus int32

const (
	ExecutionActive EncounterExecutionStatus = iota
	ExecutionCompleted
	ExecutionAbandoned
)

type EncounterExecutionDto struct {
	ID                     int64                    `json:"id"`
	EncounterId            int64                    `json:"encounterId"`
	TouristId              int64                    `json:"touristId"`
	Status                 EncounterExecutionStatus `json:"status"`
	LastActivity           time.Time                `json:"lastActivity"`
	LocationEntryTimestamp *time.Time               `json:"locationEntryTimestamp"`
	LastPosition           CoordinateDto            `json:"lastPosition"`
}
