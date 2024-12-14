package payperday

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/dto"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/utils"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure"
)

// PayPerDay represents the pay-per-day usecase implementation
type PayPerDay struct {
	infrastructure infrastructure.Infrastructure
}

// NewPayPerDay initializes the new pay-per-day implementation
func NewPayPerDay(infra infrastructure.Infrastructure) *PayPerDay {
	return &PayPerDay{
		infrastructure: infra,
	}
}

func (p *PayPerDay) CreateCoverType(ctx context.Context, coverTypeInput *dto.CoverTypeInput) (*domain.CoverType, error) {
	if !coverTypeInput.Type.IsValid() {
		return nil, fmt.Errorf("cover type must be valid. Either TPO or COMPREHENSIVE")
	}

	coverType := &domain.CoverType{
		Name: coverTypeInput.Name,
		Code: coverTypeInput.Code,
		Type: coverTypeInput.Type,
	}

	return p.infrastructure.Repository.CreateCoverType(ctx, coverType)
}

// ListCoverTypes returns a paginated collection of cover types
func (p *PayPerDay) ListCoverTypes(ctx context.Context, pagination *domain.Pagination) (*domain.CoverTypeResponse, error) {
	return p.infrastructure.Repository.ListCoverTypes(ctx, pagination)
}

// CreateProductRate is used to create a new pricing data
func (p *PayPerDay) CreateProductRate(ctx context.Context, productRateInput *dto.ProductRateInput) (*domain.ProductRate, error) {
	underwriterProduct, err := p.infrastructure.Repository.GetUnderwriterProductByID(ctx, productRateInput.ProductID)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	if !underwriterProduct.IsActive {
		return nil, fmt.Errorf("product is not active")
	}

	coverType, err := p.infrastructure.Repository.GetCoverTypeByID(ctx, productRateInput.CoverTypeID)
	if err != nil {
		return nil, fmt.Errorf("cover type not found: %w", err)
	}

	if productRateInput.Rate < 0 {
		return nil, fmt.Errorf("rate must be a non-negative number")
	}

	return p.infrastructure.Repository.CreateProductRate(ctx, &domain.ProductRate{
		Product: &domain.UnderwriterProduct{
			ID: productRateInput.ProductID,
		},
		CoverTypeID: coverType.ID,
		Rate:        productRateInput.Rate,
	})
}

// ListProductRatess returns a paginated collection of pricings
func (p *PayPerDay) ListProductRates(ctx context.Context, pagination *domain.Pagination) (*domain.ProductRateResponse, error) {
	return p.infrastructure.Repository.ListProductRates(ctx, pagination)
}

// CalculatePremiumAmount is used to calculate daily, monthly and Annual premium prices for a product
func (p *PayPerDay) CalculatePremiumAmount(ctx context.Context, input dto.PurchaseCoverInput) (*dto.PremiumAmount, error) {
	err := input.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	rate, err := p.infrastructure.Repository.GetProductRateByCoverID(ctx, input.CoverTypeID)
	if err != nil {
		return nil, fmt.Errorf("rate not found: %w", err)
	}

	floatVehicleValue, err := strconv.ParseFloat(input.VehicleValue, 64)
	if err != nil || floatVehicleValue <= 0 {
		return nil, fmt.Errorf("invalid vehicle value: %w", err)
	}

	// Actual premium (rate * vehicle value)
	actualPremium := (floatVehicleValue * rate.Rate) / 100

	// Deferred premium
	deferredPremium := (actualPremium * 115) / 100

	// Actual underwriter premium
	actualUnderwriterPremium := deferredPremium / 11

	// Round off to two decimal places
	monthlyPremium := math.Round((actualUnderwriterPremium+900)*100) / 100

	switch input.Period {
	case enums.DailyCover:
		return &dto.PremiumAmount{
			Amount: math.Round((monthlyPremium/30)*100) / 100,
		}, nil
	case enums.MonthlyCover:
		return &dto.PremiumAmount{
			Amount: monthlyPremium,
		}, nil
	case enums.AnnualCover:
		return &dto.PremiumAmount{
			Amount: deferredPremium,
		}, nil
	default:
		return nil, fmt.Errorf("invalid period: %s", input.Period)
	}
}

func (p *PayPerDay) RegisterNewUser(ctx context.Context, userPayload *dto.UserInput) (*domain.User, error) {
	err := userPayload.Validate()
	if err != nil {
		return nil, fmt.Errorf("incomplete user input: %w", err)
	}

	hashedPassword, err := utils.HashPassword(userPayload.Password)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}

	currentTime := time.Now()

	payload := &domain.User{
		Name:                                 strings.ToUpper(userPayload.Name),
		MSISDN:                               utils.FormatPhoneNumber(userPayload.MSISDN),
		IDNumber:                             userPayload.IDNumber,
		Email:                                utils.RemoveDoubleWhitespace(strings.ToLower(userPayload.Email)),
		KRAPIN:                               utils.RemoveDoubleWhitespace(userPayload.KRAPIN),
		Password:                             hashedPassword,
		IsActive:                             true,
		IsAgent:                              false,
		DateCreated:                          &currentTime,
		UpdatedAt:                            &currentTime,
		HasPaidFirstMonthlyInstallmentInFull: false,
	}

	return p.infrastructure.Repository.RegisterNewUser(ctx, payload)
}

// RegisterNewVehicle is used to register new vehicle
func (p *PayPerDay) RegisterNewVehicle(ctx context.Context, vehiclePayload *dto.VehicleInput) (*domain.VehicleInformation, error) {
	err := vehiclePayload.Validate()
	if err != nil {
		return nil, fmt.Errorf("incomplete vehicle input: %w", err)
	}

	// TODO: Get the owner from the user_id imbued in the request context
	user, err := p.infrastructure.Repository.GetUserByID(ctx, vehiclePayload.Owner)
	if err != nil {
		return nil, fmt.Errorf("the car owner does not exist as a user: %w", err)
	}

	formattedDate, err := time.Parse(time.DateOnly, vehiclePayload.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	vehicleDetails := &domain.VehicleInformation{
		ChassisNumber:      vehiclePayload.ChassisNumber,
		RegistrationNumber: vehiclePayload.RegistrationNumber,
		Make:               vehiclePayload.Make,
		Model:              vehiclePayload.Model,
		Date:               &formattedDate,
		Owner:              user.ID,
	}

	return p.infrastructure.Repository.RegisterNewVehicle(ctx, vehicleDetails)
}

// GetUserProfileByUserID returns a user profile
func (p *PayPerDay) GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := p.infrastructure.Repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		MSISDN:   user.MSISDN,
		IDNumber: user.IDNumber,
		Email:    user.Email,
		KRAPIN:   user.KRAPIN,
		IsAgent:  user.IsAgent,
		IsActive: user.IsActive,
	}, nil
}
