package infrastructure

import (
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore"
)

// Infrastructure ...
type Infrastructure struct {
	DB *datastore.DbServiceImpl
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor(
	db *datastore.DbServiceImpl,
) Infrastructure {
	return Infrastructure{
		DB: db,
	}
}
