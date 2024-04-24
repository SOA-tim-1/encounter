package handler

import (
	"database-example/dtos"
	"database-example/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type EncounterExecutionHandler struct {
	EncounterExecutionService service.IEncounterExecutionService
	EncounterService          service.IEncounterService
}

func (handler *EncounterExecutionHandler) Activate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encounterId := params["encounterId"]

	var currentPosition dtos.CoordinateDto
	err := json.NewDecoder(r.Body).Decode(&currentPosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	touristId := getPersonIdFromRequest(r)
	result, err := handler.EncounterExecutionService.Activate(encounterId, touristId, currentPosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) CompleteMiscEncounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	executionId := vars["executionId"]

	touristId := getPersonIdFromRequest(r)
	result, err := handler.EncounterExecutionService.CompleteMiscEncounter(executionId, touristId)
	if err != nil {
		// Handle error if needed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) Abandon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	executionId := vars["executionId"]

	touristId := getPersonIdFromRequest(r)
	result, _ := handler.EncounterExecutionService.Abandon(executionId, touristId)
	respondWithJSON(w, result)
}

func (handler *EncounterExecutionHandler) CheckIfCompleted(w http.ResponseWriter, r *http.Request) {
	var currentPosition dtos.CoordinateDto
	err := json.NewDecoder(r.Body).Decode(&currentPosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	executionId := vars["executionId"]

	var touristId = getPersonIdFromRequest(r)

	if execution, err := handler.EncounterExecutionService.CheckIfCompleted(executionId, touristId, currentPosition); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		respondWithJSON(w, execution)
	}
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

func getPersonIdFromRequest(r *http.Request) int64 {
	tokenString := r.Header.Get("Authorization")
	// Remove "Bearer " from token string
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	key := []byte("explorer_secret_key") // replace with your actual key

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		log.Println(err)
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		personIdStr := claims["personId"].(string)
		personId, err := strconv.ParseInt(personIdStr, 10, 64)
		if err != nil {
			log.Println("Error parsing personId:", err)
			return 0
		}
		return personId
	} else {
		log.Println("Invalid token or claims")
		return 0
	}
}
