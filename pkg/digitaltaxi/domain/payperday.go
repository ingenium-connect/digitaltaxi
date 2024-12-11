package domain

import "github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"

type CoverType struct {
	ID   string          `json:"id,omitempty"`
	Name string          `json:"name,omitempty"`
	Code string          `json:"code,omitempty"`
	Type enums.CoverType `json:"type,omitempty"`
}

type Pagination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type CoverTypeResponse struct {
	CoverTypes []*CoverType `json:"coverTypes,omitempty"`
	TotalCount int64        `json:"totalCount,omitempty"`
}
