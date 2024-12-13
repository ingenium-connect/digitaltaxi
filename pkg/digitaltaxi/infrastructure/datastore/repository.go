package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

// Create is a collection of methods to carry out create operations on the database
type Create interface {
	CreateCoverType(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error)
	CreateProductRate(ctx context.Context, productRate *domain.ProductRate) (*domain.ProductRate, error)
	RegisterNewUser(ctx context.Context, user *domain.User) (*domain.User, error)
	RegisterNewVehicle(ctx context.Context, user *domain.VehicleInformation) (*domain.VehicleInformation, error)
}

// Query hold a collection of methods to interact with the querying of any data
type Query interface {
	ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error)
	ListProductRates(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error)
	GetCoverTypeByID(ctx context.Context, id string) (*domain.CoverType, error)
	GetUnderwriterProductByID(ctx context.Context, id string) (*domain.UnderwriterProduct, error)
	GetProductRateByCoverID(ctx context.Context, id string) (*domain.ProductRate, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

// Update is a collection of methods with the ability to update any data
type Update interface {
}

type Repository interface {
	Create
	Query
	Update
}
