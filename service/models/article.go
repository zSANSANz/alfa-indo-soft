package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog -> model for Blog object
type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Author    string             `bson:"author,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Body      string             `bson:"body,omitempty"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy string             `json:"update_by" bson:"update_by"`
	UpdatedAt *time.Time         `json:"update_at" bson:"update_at"`
	DeletedBy string             `json:"deleted_by" bson:"deleted_by"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
