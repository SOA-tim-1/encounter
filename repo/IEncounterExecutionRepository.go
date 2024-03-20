package repo

import "database-example/model"

type IEncounterExecutionRepository interface {
	ICrudRepository[model.EncounterExecution]
	GetAllForTouristId(id int64) ([]*model.EncounterExecution, error)
	GetAllForEncounterId(id int64) ([]*model.EncounterExecution, error)
	GetAllActiveForEncounterId(id int64) ([]*model.EncounterExecution, error)
}
