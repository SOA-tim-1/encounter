package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type Encounter struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name                          string             `bson:"name" json:"name"`
	Description                   string             `bson:"description" json:"description"`
	Coordinates                   Coordinate         `bson:"coordinates" json:"coordinates"`
	Xp                            int32              `bson:"xp" json:"xp"`
	Status                        EncounterStatus    `bson:"status" json:"status"`
	Type                          EncounterType      `bson:"type" json:"type"`
	Range                         int32              `bson:"range" json:"range"`
	ImageUrl                      *string            `bson:"imageUrl,omitempty" json:"imageUrl"`
	MiscEncounterTask             *string            `bson:"miscEncounterTask,omitempty" json:"miscEncounterTask"`
	SocialEncounterRequiredPeople *int32             `bson:"socialEncounterRequiredPeople,omitempty" json:"socialEncounterRequiredPeople"`
	CheckpointId                  *int64             `bson:"checkpointId,omitempty" json:"checkpointId"`
	IsRequired                    *bool              `bson:"isRequired,omitempty" json:"isRequired"`
}

func (e Encounter) IsWithinRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= float64(e.Range)
}

func (e Encounter) IsWithinHiddenLocationRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= 5
}
