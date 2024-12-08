package rest

import (
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure"
)

// AcceptedContentTypes is a list of all the accepted content types
var AcceptedContentTypes = []string{"application/json", "application/x-www-form-urlencoded"}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	infrastructure infrastructure.Infrastructure
}

// NewPresentationHandlers initializes a new rest handlers usecase
func NewPresentationHandlers(infrastructure infrastructure.Infrastructure) *PresentationHandlersImpl {
	return &PresentationHandlersImpl{
		infrastructure: infrastructure,
	}
}
