package dto

import (
	"github.com/go-playground/validator"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
)

type CoverTypeInput struct {
	Name string          `json:"name"`
	Code string          `json:"code"`
	Type enums.CoverType `json:"type"`
}

type ProductRateInput struct {
	ProductID   string  `json:"product_id"`
	CoverTypeID string  `json:"covertype_id"`
	Rate        float64 `json:"rate"`
}

type PurchaseCoverInput struct {
	CoverTypeID  string            `json:"covertype_id" validate:"required"`
	VehicleValue string            `json:"vehicle_value" validate:"required"`
	Period       enums.CoverPeriod `json:"period" validate:"required"`
}

func (pc PurchaseCoverInput) Validate() error {
	v := validator.New()
	err := v.Struct(pc)

	return err
}

// UserInput is used to register a new user
type UserInput struct {
	Name     string `json:"name" validate:"required"`
	MSISDN   string `json:"msisdn" validate:"required"`
	IDNumber string `json:"id_number" validate:"required"`
	Email    string `json:"email" validate:"required"`
	KRAPIN   string `json:"kra_pin" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u UserInput) Validate() error {
	v := validator.New()
	err := v.Struct(u)

	return err
}
