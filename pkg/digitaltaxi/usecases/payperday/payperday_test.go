package payperday_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/dto"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
)

func TestPayPerDay_CreateCoverType(t *testing.T) {
	type args struct {
		ctx            context.Context
		coverTypeInput *dto.CoverTypeInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: create a cover type",
			args: args{
				ctx: context.Background(),
				coverTypeInput: &dto.CoverTypeInput{
					Name: gofakeit.BeerName(),
					Code: gofakeit.CarModel(),
					Type: enums.Comprehensive,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to create a cover type",
			args: args{
				ctx: context.Background(),
				coverTypeInput: &dto.CoverTypeInput{
					Name: gofakeit.BeerName(),
					Code: gofakeit.CarModel(),
					Type: enums.Comprehensive,
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: an invalid cover type",
			args: args{
				ctx: context.Background(),
				coverTypeInput: &dto.CoverTypeInput{
					Name: gofakeit.BeerName(),
					Code: gofakeit.CarModel(),
					Type: "invalid",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payperday, mock := setupMocks()

			if tt.name == "Sad case: unable to create a cover type" {
				mock.DataStoreMock.MockCreateCoverTypeFn = func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
					return nil, fmt.Errorf("an error occurred while creating cover type")
				}
			}
			if tt.name == "Sad case: an invalid cover type" {
				mock.DataStoreMock.MockCreateCoverTypeFn = func(ctx context.Context, coverType *domain.CoverType) (*domain.CoverType, error) {
					return nil, fmt.Errorf("an error occurred while creating cover type")
				}
			}

			_, err := payperday.CreateCoverType(tt.args.ctx, tt.args.coverTypeInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayPerDay.CreateCoverType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPayPerDay_ListCoverTypes(t *testing.T) {
	type args struct {
		ctx        context.Context
		pagination *domain.Pagination
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: list cover types",
			args: args{
				ctx: context.Background(),
				pagination: &domain.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to list cover types",
			args: args{
				ctx: context.Background(),
				pagination: &domain.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payperday, mock := setupMocks()

			if tt.name == "Sad case: unable to list cover types" {
				mock.DataStoreMock.MockListCoverTypesFn = func(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
					return nil, fmt.Errorf("an error occurred while listing cover types")
				}
			}

			_, err := payperday.ListCoverTypes(tt.args.ctx, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayPerDay.ListCoverTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPayPerDay_ListProductRates(t *testing.T) {
	type args struct {
		ctx        context.Context
		pagination *domain.Pagination
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: list rates",
			args: args{
				ctx: context.Background(),
				pagination: &domain.Pagination{
					Page:     1,
					PageSize: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to list rates",
			args: args{
				ctx: context.Background(),
				pagination: &domain.Pagination{
					Page:     1,
					PageSize: 3,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payperday, mock := setupMocks()

			if tt.name == "Sad case: unable to list rates" {
				mock.DataStoreMock.MockListProductRatesFn = func(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
					return nil, fmt.Errorf("an error occurred while listing pricings")
				}
			}

			_, err := payperday.ListProductRates(tt.args.ctx, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayPerDay.ListProductRatess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPayPerDay_CreateProductRate(t *testing.T) {
	type args struct {
		ctx              context.Context
		productRateInput *dto.ProductRateInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: create product rate",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        10,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to create product rate",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        10,
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to get an underwriter product",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        10,
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: underwriter product is inactive",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        10,
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to get cover type by ID",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        10,
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to create product rate - Negative rate",
			args: args{
				ctx: context.Background(),
				productRateInput: &dto.ProductRateInput{
					ProductID:   gofakeit.UUID(),
					CoverTypeID: gofakeit.UUID(),
					Rate:        -10,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payperday, mock := setupMocks()

			if tt.name == "Sad case: unable to create product rate" {
				mock.DataStoreMock.MockCreateProductRateFn = func(ctx context.Context, productRate *domain.ProductRate) (*domain.ProductRate, error) {
					return nil, fmt.Errorf("an error occurred while creating product rate")
				}
			}
			if tt.name == "Sad case: unable to get an underwriter product" {
				mock.DataStoreMock.MockGetUnderwriterProductByIDFn = func(ctx context.Context, id string) (*domain.UnderwriterProduct, error) {
					return nil, fmt.Errorf("an error occurred while getting underwriter product")
				}
			}
			if tt.name == "Sad case: underwriter product is inactive" {
				mock.DataStoreMock.MockGetUnderwriterProductByIDFn = func(ctx context.Context, id string) (*domain.UnderwriterProduct, error) {
					return &domain.UnderwriterProduct{IsActive: false}, nil
				}
			}
			if tt.name == "Sad case: unable to get cover type by ID" {
				mock.DataStoreMock.MockGetCoverTypeByIDFn = func(ctx context.Context, id string) (*domain.CoverType, error) {
					return nil, fmt.Errorf("an error occurred while getting cover type")
				}
			}

			_, err := payperday.CreateProductRate(tt.args.ctx, tt.args.productRateInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayPerDay.CreateProductRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPayPerDay_RegisterNewVehicle(t *testing.T) {
	type args struct {
		ctx            context.Context
		vehiclePayload *dto.VehicleInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: register new vehicle",
			args: args{
				ctx: context.Background(),
				vehiclePayload: &dto.VehicleInput{
					ChassisNumber:      "123",
					RegistrationNumber: "123",
					Make:               "Toyota",
					Model:              "Camry",
					Date:               time.Now().Format(time.DateOnly),
					Owner:              gofakeit.UUID(),
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to register new vehicle - owner is not registered",
			args: args{
				ctx: context.Background(),
				vehiclePayload: &dto.VehicleInput{
					ChassisNumber:      "123",
					RegistrationNumber: "123",
					Make:               "Toyota",
					Model:              "Camry",
					Date:               time.Now().Format(time.DateOnly),
					Owner:              gofakeit.UUID(),
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to register new vehicle",
			args: args{
				ctx: context.Background(),
				vehiclePayload: &dto.VehicleInput{
					ChassisNumber: "123",
					Make:          "Toyota",
					Model:         "Camry",
					Date:          time.Now().Format(time.DateOnly),
					Owner:         gofakeit.UUID(),
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to register new vehicle (no chassis number)",
			args: args{
				ctx: context.Background(),
				vehiclePayload: &dto.VehicleInput{
					RegistrationNumber: "123",
					Make:               "Toyota",
					Model:              "Camry",
					Date:               time.Now().Format(time.DateOnly),
					Owner:              gofakeit.UUID(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payperday, mock := setupMocks()

			if tt.name == "Sad case: unable to register new vehicle" {
				mock.DataStoreMock.MockRegisterNewVehicleFn = func(ctx context.Context, vehicleInformation *domain.VehicleInformation) (*domain.VehicleInformation, error) {
					return nil, fmt.Errorf("an error occurred while registering new vehicle")
				}
			}
			if tt.name == "Sad case: unable to register new vehicle - owner is not registered" {
				mock.DataStoreMock.MockGetUserByIDFn = func(ctx context.Context, id string) (*domain.User, error) {
					return nil, fmt.Errorf("failed to register new vehicle. Owner is not registered")
				}
			}

			_, err := payperday.RegisterNewVehicle(tt.args.ctx, tt.args.vehiclePayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayPerDay.RegisterNewVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
