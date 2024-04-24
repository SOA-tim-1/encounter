package service

import "database-example/dtos"

type IEncounterExecutionService interface {
	Activate(encounterId string, touristId int64, currentPosition dtos.CoordinateDto) (*dtos.EncounterExecutionDto, error)
	Abandon(executionId string, touristId int64) (*dtos.EncounterExecutionDto, error)
	CheckIfCompleted(executionId string, touristId int64, currentPosition dtos.CoordinateDto) (*dtos.EncounterExecutionDto, error)
	CompleteMiscEncounter(executionId string, touristId int64) (*dtos.EncounterExecutionDto, error)
}
