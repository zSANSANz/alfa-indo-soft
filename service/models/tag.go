package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID             primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
	TermID         uint               `json:"term_id" bson:"term_id"`
	TagSlug        string             `json:"tag_slug" bson:"tag_slug"`
	TagName        string             `json:"tag_name" bson:"tag_name"`
	TagDescription string             `json:"tag_description" bson:"tag_description"`
	TagCount       string             `json:"tag_count" bson:"tag_count"`
	CreatedBy      string             `json:"created_by" bson:"created_by"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy      string             `json:"update_by" bson:"update_by"`
	UpdatedAt      *time.Time         `json:"update_at" bson:"update_at"`
	DeletedBy      string             `json:"deleted_by" bson:"deleted_by"`
	DeletedAt      *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
