package service_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id"`
	NumberOfGuests *int               `json:"number-of-guests" validate:"required"`
	TableNumber    *int               `json:"number-of-table" validate:"required"`
	CreatedAt      time.Time          `json:"created-at"`
	UpdatedAt      time.Time          `json:"updated-at"`
	TableID        string             `json:"table-id"`
}
