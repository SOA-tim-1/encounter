package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterExecutionRepository) Get(id string) (model.EncounterExecution, error) {
	encounterExecution := model.EncounterExecution{}
	dbResult := repo.DatabaseConnection.First(&encounterExecution, "\"Id\" = ?", id)
	if dbResult.Error != nil {
		return model.EncounterExecution{}, dbResult.Error
	}
	return encounterExecution, nil
}

func (repo *EncounterExecutionRepository) GetAll() ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Find(&encounterExecutions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return encounterExecutions, nil
}

func (repo *EncounterExecutionRepository) Create(entity *model.EncounterExecution) (model.EncounterExecution, error) {
	dbResult := repo.DatabaseConnection.Create(entity)
	if dbResult.Error != nil {
		return model.EncounterExecution{}, dbResult.Error
	}

	return *entity, nil
}

func (repo *EncounterExecutionRepository) Update(entity *model.EncounterExecution) (model.EncounterExecution, error) {
	dbResult := repo.DatabaseConnection.Save(entity)
	if dbResult.Error != nil {
		return model.EncounterExecution{}, dbResult.Error
	}

	return *entity, nil
}

func (repo *EncounterExecutionRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.EncounterExecution{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncounterExecutionRepository) GetAllForEncounterId(id string) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	result := repo.DatabaseConnection.Where("\"EncounterId\" = ?", id).Find(&encounterExecutions)
	if result.Error != nil {
		return nil, result.Error
	}
	return encounterExecutions, nil
}

func (repo *EncounterExecutionRepository) GetAllActiveForEncounterId(id string) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	result := repo.DatabaseConnection.Where("\"EncounterId\" = ? AND \"Status\" = ?", id, model.ExecutionActive).Find(&encounterExecutions)
	if result.Error != nil {
		return nil, result.Error
	}
	return encounterExecutions, nil
}

func (repo *EncounterExecutionRepository) GetAllForTouristId(id int64) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	result := repo.DatabaseConnection.Where("\"TouristId\" = ?", id).Find(&encounterExecutions)
	if result.Error != nil {
		return nil, result.Error
	}
	return encounterExecutions, nil
}
