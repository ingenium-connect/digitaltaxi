package domain

import "time"

type VehicleInformation struct {
	ID                 string     `json:"id,omitempty"`
	ChassisNumber      string     `json:"chassis_number,omitempty"`
	RegistrationNumber string     `json:"registration_number,omitempty"`
	Make               string     `json:"make,omitempty"`
	Model              string     `json:"model,omitempty"`
	Date               *time.Time `json:"date,omitempty"`
	Owner              string     `json:"owner,omitempty"`
}
