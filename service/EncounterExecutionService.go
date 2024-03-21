package service

import (
	"database-example/dtos"
	"database-example/model"
	"database-example/repo"
	"errors"
	"log"

	"github.com/rafiulgits/go-automapper"
)

type ErrorType int

type EncounterExecutionService struct {
	EncounterExecutionRepo repo.IEncounterExecutionRepository
	EncounterService       IEncounterService
}

func (service *EncounterExecutionService) Abandon(executionId, touristId int64) (executionDto *dtos.EncounterExecutionDto, err error) {
	var encounterExecution model.EncounterExecution

	log.Println("Abandoning encounter execution with id: ", executionId)

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
	automapper.Map(encounterExecution, executionDto)
	executionDto.Status = dtos.EncounterExecutionStatus(encounterExecution.Status)
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
	encounter.Status = model.EncounterStatus(encounterDto.Status)
	encounter.Type = model.EncounterType(encounterDto.Type)
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

	var executionDto dtos.EncounterExecutionDto
	automapper.Map(newExecution, &executionDto)
	executionDto.Status = dtos.EncounterExecutionStatus(newExecution.Status)
	return &executionDto, nil
}

func (service *EncounterExecutionService) CheckIfCompleted(executionId, touristId int64, currentPositionDto dtos.CoordinateDto) (executionDto *dtos.EncounterExecutionDto, err error) {
	var encounterExecution = &model.EncounterExecution{}
	if *encounterExecution, err = service.EncounterExecutionRepo.Get(executionId); err != nil {
		return
	}

	if encounterExecution.TouristId != touristId {
		err = errors.New("forbidden")
		return
	}

	var encounterDto *dtos.EncounterDto
	if encounterDto, err = service.EncounterService.Get(encounterExecution.EncounterId); err != nil {
		return
	}

	var currentPosition model.Coordinate
	var encounter model.Encounter
	automapper.Map(currentPositionDto, &currentPosition)
	automapper.Map(encounterDto, &encounter)
	encounter.Type = model.EncounterType(encounterDto.Type)
	encounter.Status = model.EncounterStatus(encounterDto.Status)

	switch encounterDto.Type {
	case dtos.HiddenLocation:
		if err = encounterExecution.CheckIfCompletedHiddenLocation(encounter, currentPosition); err != nil {
			return
		}
		if *encounterExecution, err = service.EncounterExecutionRepo.Update(encounterExecution); err != nil {
			return
		}
	case dtos.Social:
		if encounterExecution, err = service.checkIfCompletedSocial(encounterExecution, &encounter, &currentPosition); err != nil {
			return
		}
	case dtos.Misc:
		encounterExecution.UpdateLastActivityInformation(&currentPosition)
		if *encounterExecution, err = service.EncounterExecutionRepo.Update(encounterExecution); err != nil {
			return
		}
	default:
		err = errors.New("encounter type not supported")
		return
	}

	executionDto = &dtos.EncounterExecutionDto{}
	automapper.Map(encounterExecution, executionDto)
	executionDto.Status = dtos.EncounterExecutionStatus(encounterExecution.Status)
	return executionDto, nil
}

func (service *EncounterExecutionService) checkIfCompletedSocial(encounterExecution *model.EncounterExecution, encounter *model.Encounter,
	currentPosition *model.Coordinate) (execution *model.EncounterExecution, err error) {
	if !encounter.IsWithinRange(*currentPosition) {
		encounterExecution.Abandon()
		if *encounterExecution, err = service.EncounterExecutionRepo.Update(encounterExecution); err != nil {
			return
		}
		return encounterExecution, nil
	}
	activeExecutions, _ := service.EncounterExecutionRepo.GetAllActiveForEncounterId(encounter.ID)
	if int32(len(activeExecutions)) >= *encounter.SocialEncounterRequiredPeople {
		for _, ee := range activeExecutions {
			ee.Complete(currentPosition)
			if _, err = service.EncounterExecutionRepo.Update(&ee); err != nil {
				return
			}
		}
	}
	if *encounterExecution, err = service.EncounterExecutionRepo.Get(encounterExecution.ID); err != nil {
		return nil, err
	}
	return encounterExecution, nil
}

func (service *EncounterExecutionService) CompleteMiscEncounter(executionId, touristId int64) (executionDto *dtos.EncounterExecutionDto, err error) {
	var encounterExecution model.EncounterExecution
	if encounterExecution, err = service.EncounterExecutionRepo.Get(executionId); err != nil {
		return
	}

	if encounterExecution.TouristId != touristId {
		err = errors.New("forbidden")
		return
	}

	if err = encounterExecution.Complete(nil); err != nil {
		return
	}

	if encounterExecution, err = service.EncounterExecutionRepo.Update(&encounterExecution); err != nil {
		return
	}

	executionDto = &dtos.EncounterExecutionDto{}
	automapper.Map(encounterExecution, executionDto)
	executionDto.Status = dtos.EncounterExecutionStatus(encounterExecution.Status)
	return executionDto, nil
}
