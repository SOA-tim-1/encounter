package dtos

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
	ID                            int64           `json:"id"`
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
