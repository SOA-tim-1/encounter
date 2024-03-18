package model

import (
	"math"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (c Coordinate) DistanceTo(otherCoordinate Coordinate) float64 {
	earthRadius := 6371000.0
	lat1Rad := toRadians(c.Latitude)
	lat2Rad := toRadians(otherCoordinate.Latitude)
	deltaLat := lat2Rad - lat1Rad
	deltaLon := toRadians(otherCoordinate.Longitude - c.Longitude)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	cValue := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * cValue

	return distance
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (c Coordinate) IsCloseTo(otherCoordinate Coordinate, distanceInMeters float64) bool {
	distance := c.DistanceTo(otherCoordinate)
	return distance <= distanceInMeters
}
