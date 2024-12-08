package datastore

import (
	"context"
	"time"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mongodb"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Repository holds a set of all the repository methods that interact with the database
type Repository interface {
	Create
	Query
	Update
}

// DbServiceImpl is an implementation of the database repository
// It is implementation agnostic i.e logic should be handled using
// the preferred database
type DbServiceImpl struct {
	Repository
}

// NewDbService creates a new database service
func NewDbService() *DbServiceImpl {
	ctx := context.Background()

	// This implementation is database agnostic. It can be changed to use any database. e.g. Firebase, MongoDB, etc
	environment := helpers.MustGetEnvVar("REPOSITORY")

	switch environment {
	case "mongodb":
		options := &options.ClientOptions{}
		options.ApplyURI(helpers.MustGetEnvVar("MONGODB_URI"))

		client, err := mongo.Connect(ctx, options)
		if err != nil {
			log.Panicf("can't initialize mongodb when setting up profile service: %s", err)
		}

		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		client.Ping(ctx, readpref.Primary())

		repo := mongodb.NewMongoDBClient(client.Database(helpers.MustGetEnvVar("DATABASE_NAME")))

		return &DbServiceImpl{
			Repository: repo,
		}

	case "postgres":
		return &DbServiceImpl{
			Repository: nil,
		}

	default:
		log.Panicf("unknown repository: %s", environment)
	}

	return nil
}
