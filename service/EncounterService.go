package service

import (
	"database-example/dtos"
	"database-example/repo"
)

type EncounterService struct {
	EncounterRepo repo.IEncounterRepository
}

func (service *EncounterService) Get(id string) (*dtos.EncounterDto, error) {
	encounter, err := service.EncounterRepo.Get(id)
	if err != nil {
		return nil, err
	}

	encounterDto := dtos.MapEncounterToDto(encounter)
	return &encounterDto, nil
}

func (service *EncounterService) GetAll() ([]dtos.EncounterDto, error) {
	encounters, err := service.EncounterRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var encounterDtos []dtos.EncounterDto
	for _, encounter := range encounters {
		encounterDto := dtos.MapEncounterToDto(encounter)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}

func (service *EncounterService) Create(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	encounter, err := dtos.MapDtoToEncounter(*encounterDto)
	if err != nil {
		return nil, err
	}

	createdEncounter, err := service.EncounterRepo.Create(&encounter)
	if err != nil {
		return nil, err
	}

	createdEncounterDto := dtos.MapEncounterToDto(createdEncounter)
	return &createdEncounterDto, nil
}

func (service *EncounterService) Update(encounterDto *dtos.EncounterDto) (*dtos.EncounterDto, error) {
	encounter, err := dtos.MapDtoToEncounter(*encounterDto)

	updatedEncounter, err := service.EncounterRepo.Update(&encounter)
	if err != nil {
		return nil, err
	}

	updatedEncounterDto := dtos.MapEncounterToDto(updatedEncounter)
	return &updatedEncounterDto, nil
}

func (service *EncounterService) Delete(id string) error {
	return service.EncounterRepo.Delete(id)
}

func (service *EncounterService) GetAllActive() ([]dtos.EncounterDto, error) {
	encounters, err := service.EncounterRepo.GetAllActive()
	if err != nil {
		return nil, err
	}

	var encounterDtos []dtos.EncounterDto
	for _, encounter := range encounters {
		encounterDto := dtos.MapEncounterToDto(encounter)
		encounterDtos = append(encounterDtos, encounterDto)
	}

	return encounterDtos, nil
}
