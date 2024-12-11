package mongodb

import (
	"context"
	"log"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
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

// CreateCoverTypes is used to create different types of cover types
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

// ListCoverTypes uses to retrieve all cover types
func (m *MongoDBClientImpl) ListCoverTypes(ctx context.Context, collectionName string, pagination *domain.Pagination) (*CoverTypeResponse, error) {
	totalCount, err := m.client.Collection(collectionName).CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	skip := (pagination.Page - 1) * pagination.PageSize

	pipeline := bson.A{
		bson.M{
			"$skip": int64(skip),
		}, bson.M{
			"$limit": int64(pagination.PageSize),
		},
	}

	cursor, err := m.client.Collection(collectionName).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var coverTypes []*CoverType
	if err := cursor.All(ctx, &coverTypes); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &CoverTypeResponse{
		CoverTypes: coverTypes,
		TotalCount: totalCount,
	}, nil
}
