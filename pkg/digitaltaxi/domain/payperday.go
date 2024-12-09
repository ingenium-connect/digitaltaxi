package domain

import "github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"

type CoverType struct {
	ID   string          `json:"id,omitempty"`
	Name string          `json:"name"`
	Code string          `json:"code"`
	Type enums.CoverType `json:"type"`
}
