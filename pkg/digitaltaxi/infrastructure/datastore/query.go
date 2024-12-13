package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

// ListCoverTypes retrieves a paginated list of cover types
func (s *DBImpl) ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
	output, err := s.MongoDB.ListCoverTypes(ctx, coverTypeCollectionName, pagination)
	if err != nil {
		return nil, err
	}

	var coverTypes []*domain.CoverType
	for _, coverType := range output.CoverTypes {
		coverTypes = append(coverTypes, &domain.CoverType{
			ID:   coverType.ID.Hex(),
			Name: coverType.Name,
			Code: coverType.Code,
			Type: coverType.Type,
		})
	}

	return &domain.CoverTypeResponse{
		CoverTypes: coverTypes,
		TotalCount: output.TotalCount,
	}, nil
}

// ListProductRates returns a price list
func (s *DBImpl) ListProductRates(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
	result, err := s.MongoDB.ListProductRates(ctx, productRateCollectionName, pagination)
	if err != nil {
		return nil, err
	}

	var productRates []*domain.ProductRate

	for _, rate := range result.Rates {
		underwriterProduct, err := s.MongoDB.GetUnderwriterProductByID(ctx, underwriterProductCollectionName, rate.ProductID.Hex())
		if err != nil {
			return nil, err
		}

		productRates = append(productRates, &domain.ProductRate{
			ID: rate.ID.Hex(),
			Product: &domain.UnderwriterProduct{
				ID:              underwriterProduct.ID.Hex(),
				Type:            underwriterProduct.Type,
				UnderwriterName: underwriterProduct.UnderwriterName,
				Name:            underwriterProduct.Name,
				Description:     underwriterProduct.Description,
				HasTonnage:      underwriterProduct.HasTonnage,
				HasSeats:        underwriterProduct.HasSeats,
				UnderwriterId:   underwriterProduct.UnderwriterId.Hex(),
				IsActive:        underwriterProduct.IsActive,
			},
			CoverTypeID: rate.CoverTypeID.Hex(),
			Rate:        rate.Rate,
		})
	}

	return &domain.ProductRateResponse{
		Rates:      productRates,
		TotalCount: result.TotalCount,
	}, nil
}

// GetCoverTypeByID retrieves a cover type by its ID
func (s *DBImpl) GetCoverTypeByID(ctx context.Context, id string) (*domain.CoverType, error) {
	coverType, err := s.MongoDB.GetCoverTypeByID(ctx, coverTypeCollectionName, id)
	if err != nil {
		return nil, err
	}

	return &domain.CoverType{
		ID:   coverType.ID.Hex(),
		Name: coverType.Name,
		Code: coverType.Code,
		Type: coverType.Type,
	}, nil
}

// GetProductByID is used to retrieve underwriter product given their ID
func (s *DBImpl) GetUnderwriterProductByID(ctx context.Context, id string) (*domain.UnderwriterProduct, error) {
	product, err := s.MongoDB.GetUnderwriterProductByID(ctx, underwriterProductCollectionName, id)
	if err != nil {
		return nil, err
	}

	return &domain.UnderwriterProduct{
		ID:              product.ID.Hex(),
		Type:            product.Type,
		UnderwriterName: product.UnderwriterName,
		Name:            product.Name,
		Description:     product.Description,
		HasTonnage:      product.HasTonnage,
		UnderwriterId:   product.UnderwriterId.Hex(),
		IsActive:        product.IsActive,
	}, nil
}

func (s *DBImpl) GetProductRateByCoverID(ctx context.Context, id string) (*domain.ProductRate, error) {
	rate, err := s.MongoDB.GetProductRateByCoverID(ctx, productRateCollectionName, id)
	if err != nil {
		return nil, err
	}

	underwriterProduct, err := s.MongoDB.GetUnderwriterProductByID(ctx, underwriterProductCollectionName, rate.ProductID.Hex())
	if err != nil {
		return nil, err
	}

	return &domain.ProductRate{
		ID: rate.ID.Hex(),
		Product: &domain.UnderwriterProduct{
			ID:              underwriterProduct.ID.Hex(),
			Type:            underwriterProduct.Type,
			UnderwriterName: underwriterProduct.UnderwriterName,
			Name:            underwriterProduct.Name,
			Description:     underwriterProduct.Description,
			HasTonnage:      underwriterProduct.HasTonnage,
			HasSeats:        underwriterProduct.HasSeats,
			UnderwriterId:   underwriterProduct.UnderwriterId.Hex(),
			IsActive:        underwriterProduct.IsActive,
		},
		CoverTypeID: rate.CoverTypeID.Hex(),
		Rate:        rate.Rate,
	}, nil
}
