package main

import (
	"context"
	"database-example/handler"
	"database-example/model"
	"database-example/proto/encounter"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	router := mux.NewRouter().StrictSlash(false)

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/encounter/", encounterHandler.GetAll).Methods("GET")
	api.HandleFunc("/encounter/active/", encounterHandler.GetAllActive).Methods("GET")
	api.HandleFunc("/encounter/", encounterHandler.Create).Methods("POST")
	api.HandleFunc("/encounter/", encounterHandler.Update).Methods("PUT")
	api.HandleFunc("/encounter/{id}", encounterHandler.Delete).Methods("DELETE")

	api.HandleFunc("/execution/{encounterId}", encounterExecutionHandler.Activate).Methods("POST")
	api.HandleFunc("/execution/{executionId}", encounterExecutionHandler.CheckIfCompleted).Methods("PATCH")
	api.HandleFunc("/execution/completeMisc/{executionId}", encounterExecutionHandler.CompleteMiscEncounter).Methods("PATCH")
	api.HandleFunc("/execution/abandon/{executionId}", encounterExecutionHandler.Abandon).Methods("PATCH")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8091",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router)))
}

func startServerGRPC(encounterHandler *handler.EncounterHandlergRPC) {
	listener, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	// Bootstrap gRPC server.
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Bootstrap gRPC service server and respond to request.
	//userHandler := handlers.UserHandler{}
	encounter.RegisterEncounterServiceServer(grpcServer, encounterHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
		println("Server starting")
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

}

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[encounter-api] ", log.LstdFlags)
	encounterRepoLogger := log.New(os.Stdout, "[encounter-repo] ", log.LstdFlags)

	encounterMongoRepo, err := repo.NewEncRepo(timeoutContext, encounterRepoLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer encounterMongoRepo.Disconnect(timeoutContext)

	encounterMongoRepo.Ping()

	encounterService := &service.EncounterService{EncounterRepo: encounterMongoRepo}
	//encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}

	encounterExecutionRepoLogger := log.New(os.Stdout, "[encounter-execution-repo] ", log.LstdFlags)
	encounterExecutionMongoRepo, err := repo.NewEncExeRepo(timeoutContext, encounterExecutionRepoLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer encounterExecutionMongoRepo.Disconnect(timeoutContext)

	encounterExecutionMongoRepo.Ping()

	// encounterExecutionService := &service.EncounterExecutionService{EncounterExecutionRepo: encounterExecutionMongoRepo, EncounterService: encounterService}
	// encounterExecutionHandler := &handler.EncounterExecutionHandler{EncounterExecutionService: encounterExecutionService, EncounterService: encounterService}

	//startServer(encounterHandler, encounterExecutionHandler)
	startServerGRPC(&handler.EncounterHandlergRPC{EncounterService: encounterService})
}
