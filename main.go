package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	server := getEnv("DATABASE_HOST", "localhost")
	port := getEnv("DATABASE_PORT", "5432")
	databaseName := getEnv("DATABASE_SCHEMA", "explorer-v1")
	schema := getEnv("DATABASE_SCHEMA_NAME", "encounters")
	user := getEnv("DATABASE_USERNAME", "postgres")
	password := getEnv("DATABASE_PASSWORD", "super")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", server, user, password, databaseName, port, schema)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	// newStudent := model.Student{
	// 	Person:     model.Person{Firstname: "John", Lastname: "Doe"},
	// 	Index:      "123456",
	// 	Major:      "Computer Science",
	// 	RandomData: model.RandomData{Years: 22},
	// }

	// // Kada upisemo studenta, GORM ce automatski prvo da kreira Osobu i upise u
	// // tabelu, a zatim Studenta, i to ce uraditi unutar iste transakcije.
	// database.Create(&newStudent)

	return database
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

func startServer(encounterHandler *handler.EncounterHandler, encounterExecutionHandler *handler.EncounterExecutionHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/administration/encounter/", encounterHandler.GetAll).Methods("GET")
	router.HandleFunc("/administration/encounter/active/", encounterHandler.GetAllActive).Methods("GET")
	router.HandleFunc("/administration/encounter/", encounterHandler.Create).Methods("POST")
	router.HandleFunc("/administration/encounter/", encounterHandler.Update).Methods("PUT")
	router.HandleFunc("/administration/encounter/{id}", encounterHandler.Delete).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8090",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router)))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	encounterRepo := &repo.EncounterRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{EncounterRepo: encounterRepo}
	encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}

	startServer(encounterHandler)
}
