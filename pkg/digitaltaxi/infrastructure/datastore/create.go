package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mongodb"
)

var (
	coverTypeCollectionName = "cover_type"
)

func (s *DBImpl) CreateCoverType(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
	payload := &mongodb.CoverType{
		Name: coverType.Name,
		Code: coverType.Code,
		Type: coverType.Type,
	}

	output, err := s.MongoDB.CreateCoverType(ctx, coverTypeCollectionName, payload)
	if err != nil {
		return nil, err
	}

	return &domain.CoverType{
		ID: output.ID.Hex(),
	}, nil
}
