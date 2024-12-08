package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient defines the methods that we need from mongo db driver
type MongoDBClient interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
}

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
func (m *MongoDBClientImpl) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.client.Collection("").FindOne(ctx, filter, opts...)
}
