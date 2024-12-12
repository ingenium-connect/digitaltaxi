package domain

import (
	"time"
)

type UnderwriterProduct struct {
	ID                  string    `json:"id"`
	Type                string    `json:"type"`
	UnderwriterName     string    `json:"underwriter_name"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	PeriodIDs           []string  `json:"-"`
	HasTonnage          bool      `json:"has_tonnage"`
	HasSeats            bool      `json:"has_seats"`
	Tonnes              []float64 `json:"tonnes"`
	NumberOfCertificate int32     `json:"number_of_certificates"`
	CertificatesIssued  bool      `json:"certificates_issued"`
	UnderwriterId       string    `json:"underwriter_id"`
	CreatedBy           string    `json:"created_by"`
	IsActive            bool      `json:"is_active"`
	Production          float64   `json:"production"`
	Invocations         int       `json:"invocations"`
	DateCreated         time.Time `json:"date_created"`
	UpdatedAt           time.Time `json:"updated_at"`
	Subtype             string    `json:"subtype"`
	PolicynumberType    string    `json:"policy_number_type"`
	FixedPolicyNumber   string    `json:"fixed_policy_number"`
	NumberofInstallment int       `json:"number_of_installment"`
	HasInstallment      bool      `json:"has_installment"`
}
