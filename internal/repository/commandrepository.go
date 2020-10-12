package repository

import (
	"context"
	"mongomanagementstudio/internal/driver"

	"go.mongodb.org/mongo-driver/bson"
)

type CommandRepository struct {
	store *driver.MongoStore
}

func NewCommandRepository(store *driver.MongoStore) (*CommandRepository, error) {
	return &CommandRepository{store}, nil
}

func (commandRepo *CommandRepository) RunCommand(ctx context.Context, command bson.M) bson.M {
	commandResult := commandRepo.store.Database.RunCommand(ctx, command)
	var document bson.M

	err := commandResult.Decode(&document)
	if err != nil {

	}
	return document
}
