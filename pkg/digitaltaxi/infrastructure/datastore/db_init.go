package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mongodb"
)

type MongoDBImplementation interface {
	CreateCoverType(ctx context.Context, collectionName string, coverType *mongodb.CoverType) (*mongodb.CoverType, error)
	ListCoverTypes(ctx context.Context, collectionName string, pagination *domain.Pagination) (*mongodb.CoverTypeResponse, error)
	CreateProductRate(ctx context.Context, collectionName string, pricing *mongodb.ProductRate) (*mongodb.ProductRate, error)
	ListProductRates(ctx context.Context, collectionName string, pagination *domain.Pagination) (*mongodb.ProductRateResponse, error)
	GetCoverTypeByID(ctx context.Context, collectionName, id string) (*mongodb.CoverType, error)
	GetUnderwriterProductByID(ctx context.Context, collectionName, id string) (*mongodb.UnderwriterProduct, error)
	GetProductRateByCoverID(ctx context.Context, collectionName, id string) (*mongodb.ProductRate, error)
	RegisterNewUser(ctx context.Context, collectionName string, user *mongodb.User) (*mongodb.User, error)
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
