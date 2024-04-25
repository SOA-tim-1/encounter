package dtos

import (
	"database-example/model"

	"github.com/rafiulgits/go-automapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncounterStatus int32

const (
	EncounterActive EncounterStatus = iota
	EncounterDraft
	EncounterArchived
)

type EncounterType int32

const (
	Social EncounterType = iota
	HiddenLocation
	Misc
)

type EncounterDto struct {
	ID                            string          `json:"id"`
	Name                          string          `json:"name"`
	Description                   string          `json:"description"`
	Coordinates                   CoordinateDto   `json:"coordinates"`
	Xp                            int32           `json:"xp"`
	Status                        EncounterStatus `json:"status"`
	Type                          EncounterType   `json:"type"`
	Range                         int32           `json:"range"`
	ImageUrl                      *string         `json:"imageUrl"`
	MiscEncounterTask             *string         `json:"miscEncounterTask"`
	SocialEncounterRequiredPeople *int32          `json:"socialEncounterRequiredPeople"`
	CheckpointId                  *int64          `json:"checkpointId"`
	IsRequired                    *bool           `json:"isRequired"`
}

func MapEncounterToDto(encounter model.Encounter) EncounterDto {
	encounterDto := EncounterDto{}
	automapper.Map(encounter, &encounterDto)
	encounterDto.ID = encounter.ID.Hex() // Convert ObjectID to string
	encounterDto.Type = EncounterType(encounter.Type)
	encounterDto.Status = EncounterStatus(encounter.Status)
	return encounterDto
}

func MapDtoToEncounter(encounterDto EncounterDto) (model.Encounter, error) {
	encounter := model.Encounter{}
	automapper.Map(&encounterDto, &encounter)
	if encounterDto.ID == "" {
		encounter.ID = primitive.NewObjectID()
	} else {

		id, err := primitive.ObjectIDFromHex(encounterDto.ID) // Convert string to ObjectID
		if err != nil {
			return model.Encounter{}, err
		}
		encounter.ID = id
	}

	return encounter, nil
}
