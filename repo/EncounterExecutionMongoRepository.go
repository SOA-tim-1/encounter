package repo

import (
	"context"
	"database-example/model"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type EncounterExecutionMongoRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func NewEncExeRepo(ctx context.Context, logger *log.Logger) (*EncounterExecutionMongoRepository, error) {
	dburi, exists := os.LookupEnv("MONGO_DB_URI")
	if !exists {
		dburi = "mongodb://root:pass@localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &EncounterExecutionMongoRepository{
		cli:    client,
		logger: logger,
	}, nil
}

func (repo *EncounterExecutionMongoRepository) Disconnect(ctx context.Context) error {
	err := repo.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *EncounterExecutionMongoRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := repo.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		repo.logger.Println(err)
	}

	// Print available databases
	databases, err := repo.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
	}
	fmt.Println(databases)
}

func (repo *EncounterExecutionMongoRepository) getCollection() *mongo.Collection {
	encounterDatabase := repo.cli.Database("encounterDB")
	executionsCollection := encounterDatabase.Collection("encounterExecutions")
	return executionsCollection
}

func (repo *EncounterExecutionMongoRepository) Get(id string) (model.EncounterExecution, error) {
	encounterExecution := model.EncounterExecution{}
	objectID, _ := primitive.ObjectIDFromHex(id) // Convert string to ObjectID
	filter := bson.M{"_id": objectID}

	err := repo.getCollection().FindOne(context.Background(), filter).Decode(&encounterExecution)
	if err != nil {
		repo.logger.Println(err)
		return model.EncounterExecution{}, err
	}

	return encounterExecution, nil
}

func (repo *EncounterExecutionMongoRepository) GetAll() ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	cursor, err := repo.getCollection().Find(context.Background(), bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &encounterExecutions)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounterExecutions, nil
}

func (repo *EncounterExecutionMongoRepository) Create(entity *model.EncounterExecution) (model.EncounterExecution, error) {
	_, err := repo.getCollection().InsertOne(context.Background(), entity)
	if err != nil {
		repo.logger.Println(err)
		return model.EncounterExecution{}, err
	}

	return *entity, nil
}

func (repo *EncounterExecutionMongoRepository) Update(entity *model.EncounterExecution) (model.EncounterExecution, error) {
	filter := bson.M{"_id": entity.ID}
	update := bson.M{"$set": entity}

	_, err := repo.getCollection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		repo.logger.Println(err)
		return model.EncounterExecution{}, err
	}

	return *entity, nil
}

func (repo *EncounterExecutionMongoRepository) Delete(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id) // Convert string to ObjectID
	filter := bson.M{"_id": objectID}

	_, err := repo.getCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}

	return nil
}

func (repo *EncounterExecutionMongoRepository) GetAllForEncounterId(id string) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	filter := bson.M{"encounterId": id}

	cursor, err := repo.getCollection().Find(context.Background(), filter)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &encounterExecutions)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounterExecutions, nil
}

func (repo *EncounterExecutionMongoRepository) GetAllActiveForEncounterId(id string) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	filter := bson.M{"encounterId": id, "status": model.ExecutionActive}

	cursor, err := repo.getCollection().Find(context.Background(), filter)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &encounterExecutions)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounterExecutions, nil
}

func (repo *EncounterExecutionMongoRepository) GetAllForTouristId(id int64) ([]model.EncounterExecution, error) {
	var encounterExecutions []model.EncounterExecution
	filter := bson.M{"touristId": id}

	cursor, err := repo.getCollection().Find(context.Background(), filter)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &encounterExecutions)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounterExecutions, nil
}
