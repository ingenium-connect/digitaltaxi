package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

// Create is a collection of methods to carry out create operations on the database
type Create interface {
	CreateCoverType(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error)
}

// Query hold a collection of methods to interact with the querying of any data
type Query interface {
	ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error)
}

// Update is a collection of methods with the ability to update any data
type Update interface {
}

type Repository interface {
	Create
	Query
	Update
}
