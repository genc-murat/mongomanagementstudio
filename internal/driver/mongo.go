package driver

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getDatabaseName(uri string) (*string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "couldnt parse mongo uri: %s", uri)
	}

	path := u.RequestURI()
	pathWithoutSlash := strings.TrimLeft(path, "/")
	s := strings.Split(pathWithoutSlash, "?")

	database := s[0]
	fmt.Println((database))
	return &database, nil
}

// MongoStore is an implementation of `Store` using MongoDB?
type MongoStore struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewMongoStore creates a connection to a mongo instance.
func NewMongoStore(ctx context.Context, uri, collectionName string) (*MongoStore, error) {
	mongoURI := uri

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI).SetMaxPoolSize(50))
	if err != nil {
		return nil, errors.Wrapf(err, "couldnt create mongo client %s", mongoURI)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "couldnt connect to %s", mongoURI)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrapf(err, "couldnt ping mongo server at %s", mongoURI)
	}

	database, err := getDatabaseName(mongoURI)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt get database name")
	}

	db := client.Database(*database)

	return &MongoStore{
		client,
		db,
	}, nil
}
