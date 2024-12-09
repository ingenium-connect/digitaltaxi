package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mongodb"
)

type MongoDBImplementation interface {
	CreateCoverType(ctx context.Context, collectionName string, coverType *mongodb.CoverType) (*mongodb.CoverType, error)
}

type DBImpl struct {
	MongoDB MongoDBImplementation
}

func NewDB(
	mongoDB MongoDBImplementation,
) *DBImpl {
	return &DBImpl{
		MongoDB: mongoDB,
	}
}
