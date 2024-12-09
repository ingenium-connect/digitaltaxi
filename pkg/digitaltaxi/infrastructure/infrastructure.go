package infrastructure

import (
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore"
)

// Infrastructure ...
type Infrastructure struct {
	Repository datastore.Repository
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor(
	repo datastore.Repository,
) Infrastructure {
	return Infrastructure{
		Repository: repo,
	}
}
