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
	Name                          string          `json:"name" gorm:"column:Name"`
	Description                   string          `json:"description" gorm:"column:Description"`
	Coordinates                   Coordinate      `json:"coordinates" gorm:"type:jsonb;column:Coordinates"`
	Xp                            int32           `json:"xp" gorm:"column:Xp"`
	Status                        EncounterStatus `json:"status" gorm:"column:Status"`
	Type                          EncounterType   `json:"type" gorm:"column:Type"`
	Range                         int32           `json:"range" gorm:"column:Range"`
	ImageUrl                      *string         `json:"imageUrl" gorm:"column:ImageUrl"`
	MiscEncounterTask             *string         `json:"miscEncounterTask" gorm:"column:MiscEncounterTask"`
	SocialEncounterRequiredPeople *int32          `json:"socialEncounterRequiredPeople" gorm:"column:SocialEncounterRequiredPeople"`
	CheckpointId                  *int64          `json:"checkpointId" gorm:"column:CheckpointId"`
	IsRequired                    *bool           `json:"isRequired" gorm:"column:IsRequired"`
}

func (Encounter) TableName() string {
	return "Encounters"
}

func (e Encounter) IsWithinRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= float64(e.Range)
}

func (e Encounter) IsWithinHiddenLocationRange(coordinate Coordinate) bool {
	return e.Coordinates.DistanceTo(coordinate) <= 5
}
