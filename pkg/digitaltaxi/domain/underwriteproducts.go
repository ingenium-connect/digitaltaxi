package domain

import (
	"time"
)

type UnderwriterProduct struct {
	ID                  string    `json:"id,omitempty"`
	Type                string    `json:"type,omitempty"`
	UnderwriterName     string    `json:"underwriter_name,omitempty"`
	Name                string    `json:"name,omitempty"`
	Description         string    `json:"description,omitempty"`
	PeriodIDs           []string  `json:"-"`
	HasTonnage          bool      `json:"has_tonnage,omitempty"`
	HasSeats            bool      `json:"has_seats,omitempty"`
	Tonnes              []float64 `json:"tonnes,omitempty"`
	NumberOfCertificate int32     `json:"number_of_certificates,omitempty"`
	CertificatesIssued  bool      `json:"certificates_issued,omitempty"`
	UnderwriterId       string    `json:"underwriter_id,omitempty"`
	CreatedBy           string    `json:"created_by,omitempty"`
	IsActive            bool      `json:"is_active,omitempty"`
	Production          float64   `json:"production,omitempty"`
	Invocations         int       `json:"invocations,omitempty"`
	DateCreated         time.Time `json:"date_created,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	Subtype             string    `json:"subtype,omitempty"`
	PolicynumberType    string    `json:"policy_number_type,omitempty"`
	FixedPolicyNumber   string    `json:"fixed_policy_number,omitempty"`
	NumberofInstallment int       `json:"number_of_installment,omitempty"`
	HasInstallment      bool      `json:"has_installment,omitempty"`
}
