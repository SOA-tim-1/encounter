package handler

import (
	"database-example/dtos"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterExecutionHandler struct {
	EncounterExecutionService service.IEncounterExecutionService
	EncounterService          service.IEncounterService
}

func (handler *EncounterExecutionHandler) Activate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encounterId, err := strconv.ParseInt(params["encounterId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var currentPosition dtos.CoordinateDto
	err = json.NewDecoder(r.Body).Decode(&currentPosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	touristId := int64(0) // Tourist Id
	result, _ := handler.EncounterExecutionService.Activate(encounterId, touristId, currentPosition)
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) CompleteMiscEncounter(w http.ResponseWriter, r *http.Request) {
	var executionId int64
	err := json.NewDecoder(r.Body).Decode(&executionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	touristId := int64(0) // Tourist Id
	result, err := handler.EncounterExecutionService.CompleteMiscEncounter(executionId, touristId)
	if err != nil {
		// Handle error if needed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) Abandon(w http.ResponseWriter, r *http.Request) {
	var executionId int64
	err := json.NewDecoder(r.Body).Decode(&executionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	touristId := int64(0) // Tourist Id
	result, _ := handler.EncounterExecutionService.Abandon(executionId, touristId)
	respondWithJSON(w, result)
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
