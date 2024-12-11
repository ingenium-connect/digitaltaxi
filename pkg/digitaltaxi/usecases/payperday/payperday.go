package payperday

import (
	"context"
	"fmt"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/dto"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
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

func (p *PayPerDay) CreateCoverType(ctx context.Context, coverTypeInput *dto.CoverTypeInput) (*domain.CoverType, error) {
	if !coverTypeInput.Type.IsValid() {
		return nil, fmt.Errorf("cover type must be valid. Either TPO or COMPREHENSIVE")
	}

	coverType := &domain.CoverType{
		Name: coverTypeInput.Name,
		Code: coverTypeInput.Code,
		Type: coverTypeInput.Type,
	}

	return p.infrastructure.Repository.CreateCoverType(ctx, coverType)
}

// ListCoverTypes returns a paginated collection of cover types
func (p *PayPerDay) ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
	return p.infrastructure.Repository.ListCoverTypes(ctx, pagination)
}
