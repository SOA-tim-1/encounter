package service

type EncounterExecutionService interface {
	Activate(encounterId, touristId int64, currentPosition EncounterCoordinateDto) EncounterExecutionDto
	Abandon(executionId, touristId int64) EncounterExecutionDto
	CheckIfCompleted(executionId, touristId int64, currentPosition EncounterCoordinateDto) EncounterExecutionDto
	CompleteMiscEncounter(executionId, touristId int64) EncounterExecutionDto
}
