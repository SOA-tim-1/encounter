package service

import "database-example/dtos"

type IEncounterService interface {
	ICrudService[dtos.EncounterDto]
	GetAllActive() ([]dtos.EncounterDto, error)
}
