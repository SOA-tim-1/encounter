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
	Entity
	Name                          string          `json:"name"`
	Description                   string          `json:"description"`
	Coordinates                   Coordinate      `json:"coordinates" gorm:"type:json"`
	Xp                            int32           `json:"xp"`
	Status                        EncounterStatus `json:"status"`
	Type                          EncounterType   `json:"type"`
	Range                         int32           `json:"range"`
	ImageUrl                      *string         `json:"imageUrl" gorm:"column:ImageUrl"`
	MiscEncounterTask             *string         `json:"miscEncounterTask" gorm:"column:MiscEncounterTask"`
	SocialEncounterRequiredPeople *int32          `json:"socialEncounterRequiredPeople" gorm:"column:SocialEncounterRequiredPeople"`
	CheckpointId                  *int64          `json:"checkpointId" gorm:"column:CheckpointId"`
	IsRequired                    *bool           `json:"isRequired" gorm:"column:IsRequired"`
}

func (e Encounter) IsWithinRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= float64(e.Range)
}

func (e Encounter) IsWithinHiddenLocationRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= 5
}
