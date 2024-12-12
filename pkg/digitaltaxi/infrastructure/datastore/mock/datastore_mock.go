package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

type DataStoreMock struct {
	MockCreateCoverTypeFn           func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error)
	MockListCoverTypesFn            func(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error)
	MockCreateProductRateFn         func(ctx context.Context, pricing *domain.ProductRate) (*domain.ProductRate, error)
	MockListProductRatesFn          func(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error)
	MockGetCoverTypeByIDFn          func(ctx context.Context, id string) (*domain.CoverType, error)
	MockGetUnderwriterProductByIDFn func(ctx context.Context, id string) (*domain.UnderwriterProduct, error)
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
		MockCreateProductRateFn: func(ctx context.Context, pricing *domain.ProductRate) (*domain.ProductRate, error) {
			return &domain.ProductRate{
				ID:          gofakeit.UUID(),
				ProductID:   gofakeit.UUID(),
				CoverTypeID: gofakeit.UUID(),
				Rate:        10.0,
			}, nil
		},
		MockListProductRatesFn: func(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
			return &domain.ProductRateResponse{
				Rates: []*domain.ProductRate{
					{
						ID:          gofakeit.UUID(),
						ProductID:   gofakeit.UUID(),
						CoverTypeID: gofakeit.UUID(),
						Rate:        10.0,
					},
					{
						ID:          gofakeit.UUID(),
						ProductID:   gofakeit.UUID(),
						CoverTypeID: gofakeit.UUID(),
						Rate:        15.0,
					},
				},
				TotalCount: 2,
			}, nil
		},
		MockGetCoverTypeByIDFn: func(ctx context.Context, id string) (*domain.CoverType, error) {
			return &domain.CoverType{
				ID:   gofakeit.UUID(),
				Name: gofakeit.BeerName(),
				Code: gofakeit.CreditCardCvv(),
				Type: enums.Comprehensive,
			}, nil
		},
		MockGetUnderwriterProductByIDFn: func(ctx context.Context, id string) (*domain.UnderwriterProduct, error) {
			return &domain.UnderwriterProduct{
				ID:              gofakeit.UUID(),
				Type:            gofakeit.BeerAlcohol(),
				UnderwriterName: gofakeit.Name(),
				Name:            gofakeit.Name(),
				Description:     gofakeit.HipsterSentence(20),
				UnderwriterId:   gofakeit.UUID(),
				IsActive:        true,
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

// CreateProductRate mocks the implementation of getting creating new product rate
func (m *DataStoreMock) CreateProductRate(ctx context.Context, rate *domain.ProductRate) (*domain.ProductRate, error) {
	return m.MockCreateProductRateFn(ctx, rate)
}

// ListProductRates mocks the implementation of getting a list of pricings
func (m *DataStoreMock) ListProductRates(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
	return m.MockListProductRatesFn(ctx, pagination)
}

// GetCoverTypeByID mocks the implementation of getting cover type by id
func (m *DataStoreMock) GetCoverTypeByID(ctx context.Context, id string) (*domain.CoverType, error) {
	return m.MockGetCoverTypeByIDFn(ctx, id)
}

// GetUnderwriterProductByID mocks the implementation of getting underwriter product by id
func (m *DataStoreMock) GetUnderwriterProductByID(ctx context.Context, id string) (*domain.UnderwriterProduct, error) {
	return m.MockGetUnderwriterProductByIDFn(ctx, id)
}
