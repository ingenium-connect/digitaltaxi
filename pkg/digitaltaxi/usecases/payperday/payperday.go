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

// CreateProductRate is used to create a new pricing data
func (p *PayPerDay) CreateProductRate(ctx context.Context, productRateInput *dto.ProductRateInput) (*domain.ProductRate, error) {
	underwriterProduct, err := p.infrastructure.Repository.GetUnderwriterProductByID(ctx, productRateInput.ProductID)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	if !underwriterProduct.IsActive {
		return nil, fmt.Errorf("product is not active")
	}

	coverType, err := p.infrastructure.Repository.GetCoverTypeByID(ctx, productRateInput.CoverTypeID)
	if err != nil {
		return nil, fmt.Errorf("cover type not found: %w", err)
	}

	if productRateInput.Rate < 0 {
		return nil, fmt.Errorf("rate must be a non-negative number")
	}

	return p.infrastructure.Repository.CreateProductRate(ctx, &domain.ProductRate{
		ProductID:   underwriterProduct.ID,
		CoverTypeID: coverType.ID,
		Rate:        productRateInput.Rate,
	})
}

// ListProductRatess returns a paginated collection of pricings
func (p *PayPerDay) ListProductRates(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
	return p.infrastructure.Repository.ListProductRates(ctx, pagination)
}
