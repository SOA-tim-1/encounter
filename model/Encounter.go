package model

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
	ID                            int64           `json:"id" gorm:"primaryKey:autoIncrement"`
	Name                          string          `json:"name"`
	Description                   string          `json:"description"`
	Coordinates                   Coordinate      `json:"coordinates" gorm:"type:json"`
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

func (e Encounter) IsWithinRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= float64(e.Range)
}

func (e Encounter) IsWithinHiddenLocationRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= 5
}
