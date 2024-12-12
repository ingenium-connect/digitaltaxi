package dto

type PremiumAmount struct {
	Amount float64 `json:"amount"`
}

type UnderwriterProduct struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	UnderwriterName string    `json:"underwriter_name"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	HasTonnage      bool      `json:"has_tonnage"`
	HasSeats        bool      `json:"has_seats"`
	Tonnes          []float64 `json:"tonnes"`
	UnderwriterId   string    `json:"underwriter_id"`
	IsActive        bool      `json:"is_active"`
	Subtype         string    `json:"subtype"`
}
