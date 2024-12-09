package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

type DataStoreMock struct {
	MockCreateCoverTypeFn func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error)
}

func NewDataStoreMock() *DataStoreMock {
	return &DataStoreMock{
		MockCreateCoverTypeFn: func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
			return &domain.CoverType{
				ID:   gofakeit.UUID(),
				Name: gofakeit.BeerName(),
				Code: gofakeit.CreditCardCvv(),
				Type: enums.Comprehensive,
			}, nil
		},
	}
}

func (m *DataStoreMock) CreateCoverType(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
	return m.MockCreateCoverTypeFn(ctx, coverType)
}
