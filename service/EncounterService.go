package service

import (
	"database-example/dtos"
	"database-example/model"
	"database-example/repo"

	"github.com/rafiulgits/go-automapper"
)

type EncounterService struct {
	EncounterRepo repo.IEncounterRepository
}

func (service *EncounterService) Get(id int64) (*dtos.EncounterDto, error) {
	encounter, err := service.EncounterRepo.Get(id)
	if err != nil {
		return nil, err
	}

	encounterDto := dtos.EncounterDto{}
	automapper.Map(encounter, &encounterDto) // Copies actual address for pointer fields!

	return &encounterDto, nil
}

func (service *EncounterService) GetAll() ([]dtos.EncounterDto, error) {
	encounters, err := service.EncounterRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var encounterDtos []dtos.EncounterDto
	for _, encounter := range encounters {
		encounterDto := dtos.EncounterDto{}
		automapper.Map(encounter, &encounterDto)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}

func (service *EncounterService) Create(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	var encounter model.Encounter
	automapper.Map(encounterDto, &encounter)

	createdEncounter, err := service.EncounterRepo.Create(&encounter)
	if err != nil {
		return nil, err
	}

	createdEncounterDto := dtos.EncounterDto{}
	automapper.Map(createdEncounter, &createdEncounterDto)

	return &createdEncounterDto, nil
}

func (service *EncounterService) Update(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	var encounter model.Encounter
	automapper.Map(encounterDto, &encounter)

	updatedEncounter, err := service.EncounterRepo.Update(&encounter)
	if err != nil {
		return nil, err
	}

	updatedEncounterDto := dtos.EncounterDto{}
	automapper.Map(updatedEncounter, &updatedEncounterDto)

	return &updatedEncounterDto, nil
}

func (service *EncounterService) Delete(id int64) error {
	return service.EncounterRepo.Delete(id)
}

func (service *EncounterService) GetAllActive() ([]dtos.EncounterDto, error) {
	encounters, err := service.EncounterRepo.GetAllActive()
	if err != nil {
		return nil, err
	}

	var encounterDtos []dtos.EncounterDto
	for _, encounter := range encounters {
		encounterDto := dtos.EncounterDto{}
		automapper.Map(encounter, &encounterDto)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}
