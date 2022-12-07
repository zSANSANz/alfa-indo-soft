package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	ID                primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
	AuthorId          uint               `json:"author_id" bson:"author_id"`
	AuthorLogin       string             `json:"author_login" bson:"author_login"`
	AuthorEmail       string             `json:"author_email" bson:"author_email"`
	AuthorDisplayName string             `json:"author_display_name" bson:"author_display_name"`
	AuthorFirstName   string             `json:"author_first_name" bson:"author_first_name"`
	AuthorLastName    string             `json:"author_last_name" bson:"author_last_name"`
	CreatedBy         string             `json:"created_by" bson:"created_by"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy         string             `json:"update_by" bson:"update_by"`
	UpdatedAt         *time.Time         `json:"update_at" bson:"update_at"`
	DeletedBy         string             `json:"deleted_by" bson:"deleted_by"`
	DeletedAt         *time.Time         `json:"deleted_at" bson:"deleted_at"`
	Status            string             `json:"status" bson:"status"`
}
