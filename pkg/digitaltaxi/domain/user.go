package domain

import (
	"time"
)

type User struct {
	ID                                   string     `json:"id,omitempty"`
	Name                                 string     `json:"name,omitempty"`
	MSISDN                               string     `json:"msisdn,omitempty"`
	IDNumber                             string     `json:"id_number,omitempty"`
	Email                                string     `json:"email,omitempty"`
	KRAPIN                               string     `json:"kra_pin,omitempty"`
	Password                             string     `json:"password,omitempty"`
	IsActive                             bool       `json:"is_active,omitempty"`
	IsAgent                              bool       `json:"is_agent,omitempty"`
	FCMKey                               string     `json:"fcm_key,omitempty"`
	DateCreated                          *time.Time `json:"date_created,omitempty"`
	UpdatedAt                            *time.Time `json:"updated_at,omitempty"`
	HasPaidFirstMonthlyInstallmentInFull bool       `json:"has_paid_first_monthly_in_full"`
}
