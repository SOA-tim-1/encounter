package handler

import "database-example/service"

type EncounterExecutionHandlergRPC struct {
	EncounterExecutionService service.IEncounterExecutionService
	EncounterService          service.IEncounterService
}
