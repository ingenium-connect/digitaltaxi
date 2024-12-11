package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

type DataStoreMock struct {
	MockCreateCoverTypeFn func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error)
	MockListCoverTypesFn  func(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error)
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
		MockListCoverTypesFn: func(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
			return &domain.CoverTypeResponse{
				CoverTypes: []*domain.CoverType{
					{
						ID:   gofakeit.UUID(),
						Name: gofakeit.BeerName(),
						Code: gofakeit.CreditCardCvv(),
						Type: enums.Comprehensive,
					},
					{
						ID:   gofakeit.UUID(),
						Name: gofakeit.BeerName(),
						Code: gofakeit.CreditCardCvv(),
						Type: enums.ThirdParty,
					},
				},
				TotalCount: 2,
			}, nil
		},
	}
}

// CreateCoverType mocks the implementation of creating a new cover type
func (m *DataStoreMock) CreateCoverType(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
	return m.MockCreateCoverTypeFn(ctx, coverType)
}

// ListCoverTypes mocks the implementation of getting a list of cover types
func (m *DataStoreMock) ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
	return m.MockListCoverTypesFn(ctx, pagination)
}
