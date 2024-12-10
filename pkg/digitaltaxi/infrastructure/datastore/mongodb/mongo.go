package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "code", Value: 1}}, // Index on the "code" field
		Options: options.Index().SetUnique(true),
	}

	if _, err := m.client.Collection(collectionName).Indexes().CreateOne(ctx, indexModel); err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	result, err := m.client.Collection(collectionName).InsertOne(ctx, coverType)
	if err != nil {
		return nil, err
	}

	return &CoverType{
		ID: result.InsertedID.(primitive.ObjectID),
	}, nil
}
