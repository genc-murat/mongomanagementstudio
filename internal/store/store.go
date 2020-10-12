package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	RunCommand(ctx context.Context, command bson.M) bson.M
	CollectionNames(ctx context.Context) ([]string, error)
	SetConnectionString(connectionString string)
}
