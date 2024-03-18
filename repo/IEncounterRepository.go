package repo

import "database-example/model"

type IEncounterRepository interface {
	ICrudRepository[model.Encounter]
	GetAllActive() ([]model.Encounter, error)
}
