package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID               primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
	TermID           uint               `json:"term_id" bson:"term_id"`
	CategoryNicename string             `json:"category_nicename" bson:"category_nicename"`
	// Category            string             `json:"category_parent" bson:"category_parent"`
	CatName             string     `json:"cat_name" bson:"cat_name"`
	CategoryDescription string     `json:"category_description" bson:"category_description"`
	CreatedBy           string     `json:"created_by" bson:"created_by"`
	CreatedAt           time.Time  `json:"created_at" bson:"created_at"`
	UpdatedBy           string     `json:"update_by" bson:"update_by"`
	UpdatedAt           *time.Time `json:"update_at" bson:"update_at"`
	DeletedBy           string     `json:"deleted_by" bson:"deleted_by"`
	DeletedAt           *time.Time `json:"deleted_at" bson:"deleted_at"`
}
