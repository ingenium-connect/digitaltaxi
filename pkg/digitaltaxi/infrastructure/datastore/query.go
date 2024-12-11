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
