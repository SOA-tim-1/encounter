package service

import "database-example/dtos"

type IEncounterExecutionService interface {
	Activate(encounterId, touristId int64, currentPosition dtos.CoordinateDto) (*dtos.EncounterExecutionDto, error)
	Abandon(executionId, touristId int64) (*dtos.EncounterExecutionDto, error)
	CheckIfCompleted(executionId, touristId int64, currentPosition dtos.CoordinateDto) (*dtos.EncounterExecutionDto, error)
	CompleteMiscEncounter(executionId, touristId int64) (*dtos.EncounterExecutionDto, error)
}
