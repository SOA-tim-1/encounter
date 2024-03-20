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

	result := handler.EncounterExecutionService.Activate(encounterId, currentPosition)
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) CheckIfCompleted(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	executionId, err := strconv.ParseInt(params["executionId"], 10, 64)
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

	result := handler.EncounterExecutionService.CheckIfCompleted(executionId, currentPosition)
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) CompleteMiscEncounter(w http.ResponseWriter, r *http.Request) {
	var executionId int64
	err := json.NewDecoder(r.Body).Decode(&executionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := handler.EncounterExecutionService.CompleteMiscEncounter(executionId)
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) Abandon(w http.ResponseWriter, r *http.Request) {
	var executionId int64
	err := json.NewDecoder(r.Body).Decode(&executionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := handler.EncounterExecutionService.Abandon(executionId)
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
