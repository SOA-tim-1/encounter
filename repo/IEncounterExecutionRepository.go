package repo

import "database-example/model"

type IEncounterExecutionRepository interface {
	ICrudRepository[model.EncounterExecution]
	GetAllForTouristId(id int64) ([]model.EncounterExecution, error)
	GetAllForEncounterId(id string) ([]model.EncounterExecution, error)
	GetAllActiveForEncounterId(id string) ([]model.EncounterExecution, error)
}
