package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBClientImpl is the implementation of MongoDBClient
type MongoDBClientImpl struct {
	client *mongo.Database
}

// NewMongoDBClient initializes a new MongoDBClient
func NewMongoDBClient(mn *mongo.Database) *MongoDBClientImpl {
	return &MongoDBClientImpl{
		client: mn,
	}
}

// FindOne finds a single document matching the filter
func (m *MongoDBClientImpl) CreateCoverType(ctx context.Context, collectionName string, coverType *CoverType) (*CoverType, error) {
	result, err := m.client.Collection(collectionName).InsertOne(ctx, coverType)
	if err != nil {
		return nil, err
	}

	return &CoverType{
		ID: result.InsertedID.(primitive.ObjectID),
	}, nil
}
