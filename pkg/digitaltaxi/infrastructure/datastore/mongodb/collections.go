package mongodb

import (
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

// Pricing is used to create a pricing collection for a daily policy
type Pricing struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id" binding:"required"`
	CoverTypeID primitive.ObjectID `json:"covertype_id" bson:"covertype_id" binding:"required"`
	Price       float64            `json:"price" bson:"price" binding:"required"`
}
