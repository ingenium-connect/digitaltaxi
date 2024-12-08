package shortcode

import (
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure"
)

// PayPerDay represents the pay-per-day usecase implementation
type PayPerDay struct {
	infrastructure infrastructure.Infrastructure
}

// NewPayPerDay initializes the new pay-per-day implementation
func NewPayPerDay(infra infrastructure.Infrastructure) *PayPerDay {
	return &PayPerDay{
		infrastructure: infra,
	}
}
