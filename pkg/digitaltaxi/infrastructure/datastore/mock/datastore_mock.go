package mock

import (
	"context"
	"time"

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
	MockGetProductRateByCoverIDFn   func(ctx context.Context, id string) (*domain.ProductRate, error)
	MockRegisterNewUserFn           func(ctx context.Context, user *domain.User) (*domain.User, error)
	MockRegisterNewVehicleFn        func(ctx context.Context, vehicleInformation *domain.VehicleInformation) (*domain.VehicleInformation, error)
	MockGetUserByIDFn               func(ctx context.Context, id string) (*domain.User, error)
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
				ID: gofakeit.UUID(),
				Product: &domain.UnderwriterProduct{
					ID: gofakeit.UUID(),
				},
				CoverTypeID: gofakeit.UUID(),
				Rate:        10.0,
			}, nil
		},
		MockListProductRatesFn: func(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
			return &domain.ProductRateResponse{
				Rates: []*domain.ProductRate{
					{
						ID: gofakeit.UUID(),
						Product: &domain.UnderwriterProduct{
							ID: gofakeit.UUID(),
						},
						CoverTypeID: gofakeit.UUID(),
						Rate:        10.0,
					},
					{
						ID: gofakeit.UUID(),
						Product: &domain.UnderwriterProduct{
							ID: gofakeit.UUID(),
						},
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
		MockGetProductRateByCoverIDFn: func(ctx context.Context, id string) (*domain.ProductRate, error) {
			return &domain.ProductRate{
				ID: gofakeit.UUID(),
				Product: &domain.UnderwriterProduct{
					ID: gofakeit.UUID(),
				},
				CoverTypeID: gofakeit.UUID(),
				Rate:        10.0,
			}, nil
		},
		MockRegisterNewUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			return &domain.User{
				ID:          gofakeit.UUID(),
				Name:        "",
				MSISDN:      "",
				IDNumber:    "",
				Email:       "",
				KRAPIN:      "",
				Password:    "",
				IsActive:    false,
				IsAgent:     false,
				FCMKey:      "",
				DateCreated: &time.Time{},
				UpdatedAt:   &time.Time{},
			}, nil
		},
		MockRegisterNewVehicleFn: func(ctx context.Context, vehicleInformation *domain.VehicleInformation) (*domain.VehicleInformation, error) {
			return &domain.VehicleInformation{
				ID:            gofakeit.UUID(),
				ChassisNumber: gofakeit.Sentence(10),
				Make:          gofakeit.BeerName(),
				Model:         gofakeit.BeerName(),
				Date:          &time.Time{},
				Owner:         gofakeit.UUID(),
			}, nil
		},
		MockGetUserByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return &domain.User{
				ID:   gofakeit.UUID(),
				Name: gofakeit.Name(),
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

// GetProductRateByCoverID mocks the implementation of getting product rate by cover type id
func (m *DataStoreMock) GetProductRateByCoverID(ctx context.Context, id string) (*domain.ProductRate, error) {
	return m.MockGetProductRateByCoverIDFn(ctx, id)
}

// RegisterNewUser mocks the implementation of registering new user
func (m *DataStoreMock) RegisterNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return m.MockRegisterNewUserFn(ctx, user)
}

// RegisterNewVehicle mocks the implementation of registering new vehicle
func (m *DataStoreMock) RegisterNewVehicle(ctx context.Context, vehicleInformation *domain.VehicleInformation) (*domain.VehicleInformation, error) {
	return m.MockRegisterNewVehicleFn(ctx, vehicleInformation)
}

// MockGetUserByID mocks the implementation of getting user information
func (m *DataStoreMock) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return m.MockGetUserByIDFn(ctx, id)
}
