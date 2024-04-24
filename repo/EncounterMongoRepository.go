package repo

import (
	"context"
	"database-example/model"
	"errors"
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

type EncounterMongoRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func NewEncRepo(ctx context.Context, logger *log.Logger) (*EncounterMongoRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &EncounterMongoRepository{
		cli:    client,
		logger: logger,
	}, nil
}

func (repo *EncounterMongoRepository) Disconnect(ctx context.Context) error {
	err := repo.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *EncounterMongoRepository) Ping() {
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

func (repo *EncounterMongoRepository) getCollection() *mongo.Collection {
	encounterDatabase := repo.cli.Database("encounterDB")
	encountersCollection := encounterDatabase.Collection("encounters")
	return encountersCollection
}

func (repo *EncounterMongoRepository) Get(id string) (model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.logger.Println(err)
		return model.Encounter{}, err
	}

	var encounter model.Encounter
	err = encountersCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&encounter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Encounter{}, nil
		}
		repo.logger.Println(err)
		return model.Encounter{}, err
	}

	return encounter, nil
}

func (repo *EncounterMongoRepository) GetAll() ([]model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	var encounters []model.Encounter
	cursor, err := encountersCollection.Find(ctx, bson.M{"checkpointId": bson.M{"$exists": false}})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = cursor.All(ctx, &encounters); err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounters, nil
}

func (repo *EncounterMongoRepository) Create(entity *model.Encounter) (model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	result, err := encountersCollection.InsertOne(ctx, entity)
	if err != nil {
		repo.logger.Println(err)
		return model.Encounter{}, err
	}

	entity.ID = result.InsertedID.(primitive.ObjectID)
	return *entity, nil
}

func (repo *EncounterMongoRepository) Update(entity *model.Encounter) (model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	updateResult, err := encountersCollection.ReplaceOne(ctx, bson.M{"_id": entity.ID}, entity)
	if err != nil {
		repo.logger.Println(err)
		return model.Encounter{}, err
	}

	if updateResult.MatchedCount == 0 {
		return model.Encounter{}, errors.New("no document found to update")
	}

	return *entity, nil
}

func (repo *EncounterMongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.logger.Println(err)
		return err
	}

	deleteResult, err := encountersCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		repo.logger.Println(err)
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New("no document found to delete")
	}

	return nil
}

func (repo *EncounterMongoRepository) GetAllActive() ([]model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	var encounters []model.Encounter
	cursor, err := encountersCollection.Find(ctx, bson.M{"status": model.EncounterActive, "checkpointId": bson.M{"$exists": false}})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = cursor.All(ctx, &encounters); err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	return encounters, nil
}
