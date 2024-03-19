package service

import "database-example/dtos"

type IEncouterService interface {
	ICrudService[dtos.EncounterDto]
	GetAllActive() ([]dtos.EncounterDto, error)
}
