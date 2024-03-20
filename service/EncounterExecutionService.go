package service

import (
	"database-example/dtos"
	"database-example/model"
	"database-example/repo"
	"errors"

	"github.com/rafiulgits/go-automapper"
)

type ErrorType int

type EncounterExecutionService struct {
	EncounterExecutionRepo repo.IEncounterExecutionRepository
	EncounterService       IEncounterService
}

func (service *EncounterExecutionService) Abandon(executionId, touristId int64) (executionDto *dtos.EncounterExecutionDto, err error) {
	var encounterExecution model.EncounterExecution

	if encounterExecution, err = service.EncounterExecutionRepo.Get(executionId); err != nil {
		return // naked return, returns current values of named return values
	}

	if encounterExecution.TouristId != touristId {
		err = errors.New("forbidden")
		return
	}

	if err = encounterExecution.Abandon(); err != nil {
		return
	}

	if encounterExecution, err = service.EncounterExecutionRepo.Update(&encounterExecution); err != nil {
		return
	}

	executionDto = &dtos.EncounterExecutionDto{}
	automapper.Map(&encounterExecution, executionDto)
	return
}

func (service *EncounterExecutionService) Activate(encounterId, touristId int64, currentPositionDto dtos.CoordinateDto) (executionDto *dtos.EncounterExecutionDto, err error) {
	var encounterDto *dtos.EncounterDto

	if encounterDto, err = service.EncounterService.Get(encounterId); err != nil {
		return
	}

	if encounterDto.Status != dtos.EncounterActive {
		err = errors.New("encounter not active")
		return
	}

	if err = service.ifCanActivate(encounterDto, touristId); err != nil {
		return
	}

	var encounter model.Encounter
	automapper.Map(encounterDto, &encounter)
	var currentPosition model.Coordinate
	automapper.Map(currentPositionDto, &currentPosition)
	if executionDto, err = service.activateEncounter(&encounter, touristId, &currentPosition); err != nil {
		return
	}

	return executionDto, nil
}

func (service *EncounterExecutionService) ifCanActivate(encounter *dtos.EncounterDto, touristId int64) error {
	var touristExecutions []model.EncounterExecution
	var err error

	if touristExecutions, err = service.EncounterExecutionRepo.GetAllForTouristId(touristId); err != nil {
		return err
	}

	for _, execution := range touristExecutions {
		if execution.Status == model.ExecutionActive {
			return errors.New("another encounter already active for this tourist")
		}
		if execution.EncounterId == encounter.ID && execution.Status == model.ExecutionCompleted {
			return errors.New("encounter already completed")
		}
	}
	return nil
}

func (service *EncounterExecutionService) activateEncounter(encounter *model.Encounter, touristId int64, currentPosition *model.Coordinate) (*dtos.EncounterExecutionDto, error) {
	newExecution, err := model.NewEncounterExecution(*encounter, touristId, model.ExecutionActive, *currentPosition)
	if err != nil {
		return nil, err
	}

	if *newExecution, err = service.EncounterExecutionRepo.Create(newExecution); err != nil {
		return nil, err
	}

	var executionDto *dtos.EncounterExecutionDto
	automapper.Map(&newExecution, executionDto)
	return executionDto, nil
}

func (service *EncounterExecutionService) CheckIfCompleted(executionId, touristId int64, currentPosition dtos.CoordinateDto) (*dtos.EncounterExecutionDto, error) {
	encounterExecution, err := service.EncounterExecutionRepo.Get(executionId)
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

	encounterDto := encounterResult.Value
	switch encounterDto.Type {
	case HiddenLocation:
		encounterExecution.CheckIfCompletedHiddenLocation(service.Mapper.MapToEncounter(encounter), service.Mapper.MapToCoordinate(currentPosition))
		err := service.EncounterExecutionRepo.Update(encounterExecution)
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
		err := service.EncounterExecutionRepo.Update(encounterExecution)
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

func (service *EncounterExecutionService) checkIfCompletedSocial(encounterExecution *EncounterExecution, encounter *Encounter, currentPosition *Coordinate) (*EncounterExecution, error) {
	if !encounterDto.IsWithinRange(currentPosition) {
		encounterExecution.Abandon()
		err := service.EncounterExecutionRepo.Update(encounterExecution)
		if err != nil {
			return nil, err
		}
		return encounterExecution, nil
	}

	allActive, err := service.EncounterExecutionRepo.GetAllActiveForEncounterId(encounter.Id)
	if err != nil {
		return nil, err
	}

	if len(allActive) >= encounterDto.SocialEncounterRequiredPeople {
		for _, e := range allActive {
			e.Complete(currentPosition)
			err := service.EncounterExecutionRepo.Update(e)
			if err != nil {
				return nil, err
			}
		}
	}

	return service.EncounterExecutionRepo.Get(encounterExecution.Id)
}

func (service *EncounterExecutionService) CompleteMiscEncounter(executionId, touristId int64) (*dtos.EncounterExecutionDto, error) {
	encounterExecution, err := service.EncounterExecutionRepo.GetUntracked(executionId)
	if err != nil {
		return nil, err
	}

	if encounterExecution.TouristId != touristId {
		return nil, errors.New("forbidden")
	}

	encounterExecution.Complete()
	err = service.EncounterExecutionRepo.Update(encounterExecution)
	if err != nil {
		return nil, err
	}

	encounterResult, err := service.EncounterService.Get(encounterExecution.EncounterId)
	if err != nil {
		return nil, err
	}

	return service.Mapper.MapToDto(encounterExecution), nil
}

func (service *EncounterExecutionService) getByEncounterId(encounterId int64) ([]*EncounterExecutionDto, error) {
	encounterExecutions, err := service.EncounterExecutionRepo.GetAllForEncounterId(encounterId)
	if err != nil {
		return nil, err
	}

	executionDtos := make([]*EncounterExecutionDto, len(encounterExecutions))
	for i, execution := range encounterExecutions {
		executionDtos[i] = service.Mapper.MapToDto(execution)
	}

	return executionDtos, nil
}
