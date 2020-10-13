package repository

import (
	"context"
	"fmt"
	"mongomanagementstudio/internal/driver"

	"go.mongodb.org/mongo-driver/bson"
)

type CommandRepository struct {
	store            *driver.MongoStore
	connectionString string
}

func NewCommandRepository(store *driver.MongoStore, connectionString string) (*CommandRepository, error) {
	return &CommandRepository{store, connectionString}, nil
}

func (commandRepo *CommandRepository) RunCommand(ctx context.Context, command bson.M) bson.M {
	fmt.Println(command)
	commandResult := commandRepo.store.Database.RunCommand(ctx, command)
	var document bson.M

	err := commandResult.Decode(&document)
	if err != nil {

	}
	return document
}

func (commandRepo *CommandRepository) CollectionNames(ctx context.Context) ([]string, error) {
	if commandRepo.store == nil {
		commandRepo.store, _ = driver.NewMongoStore(context.Background(), commandRepo.connectionString, "")
		fmt.Println(commandRepo.connectionString)
	}
	return commandRepo.store.Database.ListCollectionNames(ctx, bson.D{})
}

func (commandRepo *CommandRepository) SetConnectionString(connectionString string) {
	commandRepo.connectionString = connectionString
}
