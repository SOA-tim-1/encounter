package dtos

import (
	"database-example/model"
	"time"

	"github.com/rafiulgits/go-automapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncounterExecutionStatus int32

const (
	ExecutionActive EncounterExecutionStatus = iota
	ExecutionCompleted
	ExecutionAbandoned
)

type EncounterExecutionDto struct {
	ID                     string                   `json:"id"`
	EncounterId            string                   `json:"encounterId"`
	TouristId              int64                    `json:"touristId"`
	Status                 EncounterExecutionStatus `json:"status"`
	LastActivity           time.Time                `json:"lastActivity"`
	LocationEntryTimestamp *time.Time               `json:"locationEntryTimestamp"`
	LastPosition           CoordinateDto            `json:"lastPosition"`
}

func MapEncounterExecutionToDto(encounterExecution model.EncounterExecution) EncounterExecutionDto {
	encounterExecutionDto := EncounterExecutionDto{}
	automapper.Map(encounterExecution, &encounterExecutionDto)
	encounterExecutionDto.ID = encounterExecution.ID.Hex()                   // Convert ObjectID to string
	encounterExecutionDto.EncounterId = encounterExecution.EncounterId.Hex() // Convert ObjectID to string
	encounterExecutionDto.Status = EncounterExecutionStatus(encounterExecution.Status)
	return encounterExecutionDto
}

func MapDtoToEncounterExecution(encounterExecutionDto EncounterExecutionDto) (model.EncounterExecution, error) {
	encounterExecution := model.EncounterExecution{}
	automapper.Map(&encounterExecutionDto, &encounterExecution)
	if encounterExecutionDto.ID == "" {
		encounterExecution.ID = primitive.NewObjectID()
	} else {
		id, err := primitive.ObjectIDFromHex(encounterExecutionDto.ID) // Convert string to ObjectID
		if err != nil {
			return model.EncounterExecution{}, err
		}
		encounterExecution.ID = id
	}

	encounterId, err := primitive.ObjectIDFromHex(encounterExecutionDto.EncounterId) // Convert string to ObjectID
	if err != nil {
		return model.EncounterExecution{}, err
	}
	encounterExecution.EncounterId = encounterId

	return encounterExecution, nil
}
