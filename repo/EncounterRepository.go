package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRepository) Get(id int64) (model.Encounter, error) {
	encounter := model.Encounter{}
	dbResult := repo.DatabaseConnection.First(&encounter, "id = ?", id)
	if dbResult.Error != nil {
		return model.Encounter{}, dbResult.Error
	}

	return encounter, nil
}

func (repo *EncounterRepository) GetAll() ([]model.Encounter, error) {
	encounters := []model.Encounter{}
	dbResult := repo.DatabaseConnection.Where("checkpointId IS NULL").Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncounterRepository) Create(entity *model.Encounter) (model.Encounter, error) {
	dbResult := repo.DatabaseConnection.Create(entity)
	if dbResult.Error != nil {
		return model.Encounter{}, dbResult.Error
	}

	return *entity, nil
}

func (repo *EncounterRepository) Update(entity *model.Encounter) (model.Encounter, error) {
	dbResult := repo.DatabaseConnection.Save(entity)
	if dbResult.Error != nil {
		return model.Encounter{}, dbResult.Error
	}

	return *entity, nil
}

func (repo *EncounterRepository) Delete(id int64) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Encounter{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

// Possibly need to adjust column names, might be lower/upper case sensitive
func (repo *EncounterRepository) GetAllActive() ([]model.Encounter, error) {
	encounters := []model.Encounter{}
	dbResult := repo.DatabaseConnection.Where("status = ? AND checkpointId IS NULL", model.EncounterActive).Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}
