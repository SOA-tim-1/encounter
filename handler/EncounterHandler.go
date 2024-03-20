package handler

import (
	"database-example/dtos"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	EncounterService service.IEncounterService
}

func (handler *EncounterHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	encounters, err := handler.EncounterService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) GetAllActive(writer http.ResponseWriter, req *http.Request) {
	encounters, err := handler.EncounterService.GetAllActive()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter dtos.EncounterDto
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdEncounter, err := handler.EncounterService.Create(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(createdEncounter)
}

func (handler *EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var encounter dtos.EncounterDto
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedEncounter, err := handler.EncounterService.Update(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedEncounter)
}

func (handler *EncounterHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	encounterId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.Delete(encounterId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
