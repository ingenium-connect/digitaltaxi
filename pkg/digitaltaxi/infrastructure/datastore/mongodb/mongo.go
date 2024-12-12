package mongodb

import (
	"context"
	"errors"
	"fmt"
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

// CreateProductRate creates a new product rate
func (m *MongoDBClientImpl) CreateProductRate(ctx context.Context, collectionName string, productRate *ProductRate) (*ProductRate, error) {
	filter := bson.M{
		"product_id":   productRate.ProductID,
		"covertype_id": productRate.CoverTypeID,
	}

	output := m.client.Collection(collectionName).FindOne(ctx, filter)
	if errors.Is(output.Err(), mongo.ErrNoDocuments) {
		result, err := m.client.Collection(collectionName).InsertOne(ctx, productRate)
		if err != nil {
			return nil, err
		}

		return &ProductRate{
			ID: result.InsertedID.(primitive.ObjectID),
		}, nil
	}

	var rate *ProductRate
	if err := output.Decode(&rate); err != nil {
		return nil, fmt.Errorf("error decoding product rate: %w", err)
	}

	return rate, nil
}

func (m *MongoDBClientImpl) ListProductRates(ctx context.Context, collectionName string, pagination *domain.Pagination) (*ProductRateResponse, error) {
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

	var rates []*ProductRate
	if err := cursor.All(ctx, &rates); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &ProductRateResponse{
		Rates:      rates,
		TotalCount: totalCount,
	}, nil
}

// GetCoverTypeByID retrieves a cover type by its ID
func (m *MongoDBClientImpl) GetCoverTypeByID(ctx context.Context, collectionName string, id string) (*CoverType, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	coverType := &CoverType{}

	err = m.client.Collection(collectionName).FindOne(ctx, bson.M{"_id": objID}).Decode(coverType)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("no documents found")
	} else if err != nil {
		return nil, err
	}

	return coverType, nil
}

// GetUnderwriterProductByID is used to retrieve underwriter product
func (m *MongoDBClientImpl) GetUnderwriterProductByID(ctx context.Context, collectionName string, id string) (*UnderwriterProduct, error) {
	underwriterProductID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	product := &UnderwriterProduct{}

	err = m.client.Collection(collectionName).FindOne(ctx, bson.M{"_id": underwriterProductID}).Decode(product)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("no documents found")
	} else if err != nil {
		return nil, err
	}

	return product, nil
}
