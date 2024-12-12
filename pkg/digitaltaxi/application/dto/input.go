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

type PurchaseCover struct {
	CoverID      string `json:"cover_id" validate:"required"`
	VehicleValue string `json:"vehicle_value" validate:"required"`
}

func (pc PurchaseCover) Validate() error {
	v := validator.New()
	err := v.Struct(pc)

	return err
}
