package payperday_test

import (
	"context"
	"fmt"
	"testing"

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
