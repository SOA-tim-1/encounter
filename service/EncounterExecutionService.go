package service

import (
	"errors"
	"fmt"

	"github.com/rafiulgits/go-automapper"
)

type ErrorType int

const (
	AlreadyActive ErrorType = iota
	AlreadyCompleted
	Success
)

type EncounterExecutionService struct {
	Repository       EncounterExecutionRepository
	EncounterService EncounterService
	Mapper           automapper.Mapper
}

func NewEncounterExecutionService(repository EncounterExecutionRepository, encounterService EncounterService, mapper automapper.Mapper) *EncounterExecutionService {
	return &EncounterExecutionService{
		Repository:       repository,
		EncounterService: encounterService,
		Mapper:           mapper,
	}
}

func (service *EncounterExecutionService) Abandon(executionId, touristId int64) (*EncounterExecutionDto, error) {
	encounterExecution, err := service.Repository.GetUntracked(executionId)
	if err != nil {
		return nil, err
	}

	if encounterExecution.TouristId != touristId {
		return nil, errors.New("forbidden")
	}

	encounterExecution.Abandon()
	err = service.Repository.Update(encounterExecution)
	if err != nil {
		return nil, err
	}

	return service.Mapper.MapToDto(encounterExecution), nil
}

func (service *EncounterExecutionService) Activate(encounterId, touristId int64, currentPosition EncounterCoordinateDto) (*EncounterExecutionDto, error) {
	result, err := service.EncounterService.Get(encounterId)
	if err != nil {
		return nil, err
	}

	encounter := result.Value
	if encounter.Status != Active {
		return nil, errors.New("encounter not active")
	}

	status := service.IfCanActivate(encounter, touristId)
	if status != Success {
		return nil, service.ReturnMatchingError(status)
	}

	execution, err := service.ActivateEncounter(service.Mapper.MapToEncounter(encounter), touristId, service.Mapper.MapToCoordinate(currentPosition))
	if err != nil {
		return nil, err
	}

	return service.Mapper.MapToDto(execution), nil
}

func (service *EncounterExecutionService) ActivateEncounter(encounter *Encounter, touristId int64, currentPosition *Coordinate) (*EncounterExecution, error) {
	newExecution, err := NewEncounterExecution(encounter, touristId, Active, currentPosition)
	if err != nil {
		return nil, err
	}

	return service.Repository.Create(newExecution)
}

func (service *EncounterExecutionService) CheckIfCompleted(executionId, touristId int64, currentPosition EncounterCoordinateDto) (*EncounterExecutionDto, error) {
	encounterExecution, err := service.Repository.Get(executionId)
	if err != nil {
		return nil, err
	}

	if encounterExecution.TouristId != touristId {
		return nil, errors.New("forbidden")
	}

	encounterResult, err := service.EncounterService.Get(encounterExecution.EncounterId)
	if err != nil {
		return nil, err
	}

	encounter := encounterResult.Value
	switch encounter.Type {
	case HiddenLocation:
		encounterExecution.CheckIfCompletedHiddenLocation(service.Mapper.MapToEncounter(encounter), service.Mapper.MapToCoordinate(currentPosition))
		err := service.Repository.Update(encounterExecution)
		if err != nil {
			return nil, err
		}
	case Social:
		processedExecution, err := service.CheckIfCompletedSocial(encounterExecution, service.Mapper.MapToEncounter(encounter), service.Mapper.MapToCoordinate(currentPosition))
		if err != nil {
			return nil, err
		}
		return service.Mapper.MapToDto(processedExecution), nil
	case Misc:
		encounterExecution.UpdateLastActivityInformation(service.Mapper.MapToCoordinate(currentPosition))
		err := service.Repository.Update(encounterExecution)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("internal server error")
	}

	if encounterExecution.Status == Completed {
	}

	return service.Mapper.MapToDto(encounterExecution), nil
}

func (service *EncounterExecutionService) CheckIfCompletedSocial(encounterExecution *EncounterExecution, encounter *Encounter, currentPosition *Coordinate) (*EncounterExecution, error) {
	if !encounter.IsWithinRange(currentPosition) {
		encounterExecution.Abandon()
		err := service.Repository.Update(encounterExecution)
		if err != nil {
			return nil, err
		}
		return encounterExecution, nil
	}

	allActive, err := service.Repository.GetAllActiveForEncounterId(encounter.Id)
	if err != nil {
		return nil, err
	}

	if len(allActive) >= encounter.SocialEncounterRequiredPeople {
		for _, e := range allActive {
			e.Complete(currentPosition)
			err := service.Repository.Update(e)
			if err != nil {
				return nil, err
			}
		}
	}

	return service.Repository.Get(encounterExecution.Id)
}

func (service *EncounterExecutionService) CompleteMiscEncounter(executionId, touristId int64) (*EncounterExecutionDto, error) {
	encounterExecution, err := service.Repository.GetUntracked(executionId)
	if err != nil {
		return nil, err
	}

	if encounterExecution.TouristId != touristId {
		return nil, errors.New("forbidden")
	}

	encounterExecution.Complete()
	err = service.Repository.Update(encounterExecution)
	if err != nil {
		return nil, err
	}

	encounterResult, err := service.EncounterService.Get(encounterExecution.EncounterId)
	if err != nil {
		return nil, err
	}

	return service.Mapper.MapToDto(encounterExecution), nil
}

func (service *EncounterExecutionService) IfCanActivate(encounter *Encounter, touristId int64) ErrorType {
	touristExecutions, err := service.Repository.GetAllForTouristId(touristId)
	if err != nil {
		return Internal
	}

	for _, enc := range touristExecutions {
		if enc.Status == Active {
			return AlreadyActive
		}
		if enc.EncounterId == encounter.Id && enc.Status == Completed {
			return AlreadyCompleted
		}
	}

	return Success
}

func (service *EncounterExecutionService) ReturnMatchingError(status ErrorType) error {
	switch status {
	case AlreadyCompleted:
		return fmt.Errorf("encounter already completed")
	case AlreadyActive:
		return fmt.Errorf("encounter already activated")
	default:
		return fmt.Errorf("internal server error")
	}
}

func (service *EncounterExecutionService) GetByEncounterId(encounterId int64) ([]*EncounterExecutionDto, error) {
	encounterExecutions, err := service.Repository.GetAllForEncounterId(encounterId)
	if err != nil {
		return nil, err
	}

	executionDtos := make([]*EncounterExecutionDto, len(encounterExecutions))
	for i, execution := range encounterExecutions {
		executionDtos[i] = service.Mapper.MapToDto(execution)
	}

	return executionDtos, nil
}
