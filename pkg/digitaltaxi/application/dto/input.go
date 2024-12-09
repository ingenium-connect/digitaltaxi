package dto

import "github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"

type CoverTypeInput struct {
	Name string          `json:"name"`
	Code string          `json:"code"`
	Type enums.CoverType `json:"type"`
}
