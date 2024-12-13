package datastore

import (
	"context"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	coverTypeCollectionName          = "cover_type"
	productRateCollectionName        = "product_rate"
	underwriterProductCollectionName = "underwriter_products"
	usersCollectionName              = "users"
	vehicleCollectionName            = "vehicle"
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

// CreateProductRate is used to create a new pricing
func (s *DBImpl) CreateProductRate(ctx context.Context, rate *domain.ProductRate) (*domain.ProductRate, error) {
	productID, err := primitive.ObjectIDFromHex(rate.Product.ID)
	if err != nil {
		return nil, err
	}

	coverTypeID, err := primitive.ObjectIDFromHex(rate.CoverTypeID)
	if err != nil {
		return nil, err
	}

	payload := &mongodb.ProductRate{
		ProductID:   productID,
		CoverTypeID: coverTypeID,
		Rate:        rate.Rate,
	}

	output, err := s.MongoDB.CreateProductRate(ctx, productRateCollectionName, payload)
	if err != nil {
		return nil, err
	}

	return &domain.ProductRate{
		ID: output.ID.Hex(),
	}, nil
}

// RegisterNewUser is used to register a new user
func (s *DBImpl) RegisterNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	payload := &mongodb.User{
		Name:                                 user.Name,
		MSISDN:                               user.MSISDN,
		IDNumber:                             user.IDNumber,
		Email:                                user.Email,
		KRAPIN:                               user.KRAPIN,
		Password:                             user.Password,
		IsActive:                             user.IsActive,
		IsAgent:                              false,
		DateCreated:                          *user.DateCreated,
		UpdatedAt:                            *user.DateCreated,
		HasPaidFirstMonthlyInstallmentInFull: user.HasPaidFirstMonthlyInstallmentInFull,
	}

	output, err := s.MongoDB.RegisterNewUser(ctx, usersCollectionName, payload)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID: output.ID.Hex(),
	}, nil
}

// RegisterNewVehicle registers a new vehicle
func (s *DBImpl) RegisterNewVehicle(ctx context.Context, vehicleInformation *domain.VehicleInformation) (*domain.VehicleInformation, error) {
	ownerID, err := primitive.ObjectIDFromHex(vehicleInformation.Owner)
	if err != nil {
		return nil, err
	}

	payload := &mongodb.VehicleInformation{
		ChassisNumber: vehicleInformation.ChassisNumber,
		Make:          vehicleInformation.Make,
		Model:         vehicleInformation.Model,
		Date:          *vehicleInformation.Date,
		Owner:         ownerID,
	}

	output, err := s.MongoDB.RegisterNewVehicle(ctx, vehicleCollectionName, payload)
	if err != nil {
		return nil, err
	}

	return &domain.VehicleInformation{
		ID: output.ID.Hex(),
	}, nil
}
