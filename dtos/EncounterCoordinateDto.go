package dtos

import "fmt"

type EncounterCoordinateDto struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewEncounterCoordinateDto(latitude, longitude float64) *EncounterCoordinateDto {
	return &EncounterCoordinateDto{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (ecd *EncounterCoordinateDto) Equals(other *EncounterCoordinateDto) bool {
	if other == nil {
		return false
	}
	return ecd.Latitude == other.Latitude && ecd.Longitude == other.Longitude
}

func (ecd *EncounterCoordinateDto) HashCode() int {
	// Generate a hash code based on the latitude and longitude
	return fmt.Sprintf("%f:%f", ecd.Latitude, ecd.Longitude)
}
