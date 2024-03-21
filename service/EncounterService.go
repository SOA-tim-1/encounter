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
	encounterDto.Type = dtos.EncounterType(encounter.Type)
	encounterDto.Status = dtos.EncounterStatus(encounter.Status)

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
		encounterDto.Type = dtos.EncounterType(encounter.Type)
		encounterDto.Status = dtos.EncounterStatus(encounter.Status)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}

func (service *EncounterService) Create(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	var encounter model.Encounter
	automapper.Map(encounterDto, &encounter)
	encounter.Type = model.EncounterType(encounterDto.Type)
	encounter.Status = model.EncounterStatus(encounterDto.Status)

	createdEncounter, err := service.EncounterRepo.Create(&encounter)
	if err != nil {
		return nil, err
	}

	createdEncounterDto := dtos.EncounterDto{}
	automapper.Map(createdEncounter, &createdEncounterDto)
	createdEncounterDto.Type = dtos.EncounterType(createdEncounter.Type)
	createdEncounterDto.Status = dtos.EncounterStatus(createdEncounter.Status)

	return &createdEncounterDto, nil
}

func (service *EncounterService) Update(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	var encounter model.Encounter
	automapper.Map(encounterDto, &encounter)
	encounter.Type = model.EncounterType(encounterDto.Type)
	encounter.Status = model.EncounterStatus(encounterDto.Status)

	updatedEncounter, err := service.EncounterRepo.Update(&encounter)
	if err != nil {
		return nil, err
	}

	updatedEncounterDto := dtos.EncounterDto{}
	automapper.Map(updatedEncounter, &updatedEncounterDto)
	updatedEncounterDto.Type = dtos.EncounterType(updatedEncounter.Type)
	updatedEncounterDto.Status = dtos.EncounterStatus(updatedEncounter.Status)

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
		encounterDto.Type = dtos.EncounterType(encounter.Type)
		encounterDto.Status = dtos.EncounterStatus(encounter.Status)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}
