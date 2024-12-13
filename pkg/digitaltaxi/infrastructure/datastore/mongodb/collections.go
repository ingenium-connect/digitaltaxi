package mongodb

import (
	"time"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CoverType is used to model cover type collection
type CoverType struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" binding:"required"`
	Code string             `json:"code" bson:"code" binding:"required"`
	Type enums.CoverType    `json:"type" bson:"type" binding:"required"`
}

// ProductRate is used to store the product rate information
type ProductRate struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id" binding:"required"`
	CoverTypeID primitive.ObjectID `json:"covertype_id" bson:"covertype_id" binding:"required"`
	Rate        float64            `json:"rate" bson:"rate" binding:"required"`
}

// CoverTypeResponse used to return the cover types
type CoverTypeResponse struct {
	CoverTypes []*CoverType `json:"coverTypes"`
	TotalCount int64        `json:"totalCount"`
}

type ProductRateResponse struct {
	Rates      []*ProductRate `json:"rates"`
	TotalCount int64          `json:"totalCount"`
}

type UnderwriterProduct struct {
	ID                  primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Type                string               `json:"type" bson:"type" binding:"required"`
	UnderwriterName     string               `json:"underwriter_name" bson:"underwriter_name" binding:"required"`
	Name                string               `json:"name" bson:"name" binding:"required"`
	Description         string               `json:"description" bson:"description"`
	PeriodIDs           []primitive.ObjectID `json:"-" bson:"periods" binding:"required"`
	HasTonnage          bool                 `json:"has_tonnage" bson:"has_tonnage" binding:"required"`
	HasSeats            bool                 `json:"has_seats" bson:"has_seats" binding:"required"`
	Tonnes              []float64            `json:"tonnes" bson:"tonnes"`
	NumberOfCertificate int32                `json:"number_of_certificates" bson:"number_of_certificates"`
	CertificatesIssued  bool                 `json:"certificates_issued" bson:"certificates_issued"`
	UnderwriterId       primitive.ObjectID   `json:"underwriter_id" bson:"underwriter_id" binding:"required"`
	CreatedBy           primitive.ObjectID   `json:"created_by" bson:"created_by"`
	IsActive            bool                 `json:"is_active" bson:"is_active" binding:"required"`
	Production          float64              `json:"production"`
	Invocations         int                  `json:"invocations"`
	DateCreated         time.Time            `json:"date_created" bson:"date_created"`
	UpdatedAt           time.Time            `json:"updated_at" bson:"updated_at"`
	Subtype             string               `json:"subtype" bson:"subtype" binding:"required"`
	PolicynumberType    string               `json:"policy_number_type" bson:"policy_number_type" `
	FixedPolicyNumber   string               `json:"fixed_policy_number" bson:"fixed_policy_number" `
	NumberofInstallment int                  `json:"number_of_installment" bson:"number_of_installment" `
	HasInstallment      bool                 `json:"has_installment" bson:"has_installment" `
}

type User struct {
	ID                                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                                 string             `json:"name" bson:"name" binding:"required"`
	MSISDN                               string             `json:"msisdn" bson:"msisdn" binding:"required"`
	IDNumber                             string             `json:"id_number" bson:"id_number"`
	Email                                string             `json:"email" bson:"email" binding:"required"`
	KRAPIN                               string             `json:"kra_pin" bson:"kra_pin" binding:"required"`
	Password                             string             `json:"-" bson:"password" binding:"required"`
	IsActive                             bool               `json:"is_active" bson:"is_active" binding:"required"`
	IsAgent                              bool               `json:"is_agent" bson:"is_agent" binding:"required"`
	FCMKey                               string             `json:"fcm_key" bson:"fcm_key" `
	Channel                              string             `json:"channel" bson:"channel"`
	DateCreated                          time.Time          `json:"date_created" bson:"date_created"`
	UpdatedAt                            time.Time          `json:"updated_at" bson:"updated_at"`
	HasPaidFirstMonthlyInstallmentInFull bool               `json:"has_paid_first_month_in_full" bson:"has_paid_first_month_in_full" binding:"required"`
}
