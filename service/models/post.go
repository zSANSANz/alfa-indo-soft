package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID    bson.ObjectId `json:"id" xml:"id" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	// Pubdate   string             `json:"pubdate" bson:"pubdate"`
	// Pubdate string `json:"pubdate" bson:"pubdate"`
	Date string `json:"date" bson:"date"`
	// Link      string             `json:"link" bson:"link"`
	Url string `json:"url" bson:"url"`
	// CreatorId bson.ObjectId `json:"creator_id" bson:"creator_id"`
	Username    string `json:"username" bson:"username"`
	Description string `json:"description" bson:"description"`
	Content     string `json:"content" bson:"content"`
	// Except        string             `json:"except" bson:"except"`
	Except        string     `json:"excerpt" bson:"excerpt"`
	CommentStatus string     `json:"comment_status" bson:"comment_status"`
	Slug          string     `json:"slug" bson:"slug"`
	Status        string     `json:"status" bson:"status"`
	Type          string     `json:"type" bson:"type"`
	PostSurvey    PostSurvey `json:"post_survey" bson:"post_survey"`
	// PostParent    string     `json:"post_parent" bson:"post_parent"`
	PostType string `json:"post_type" bson:"post_type"`
	// CategoryId    bson.ObjectId `json:"category_id" bson:"category_id"`
	// Category      []DataCategory     `bson:"category,omitempty"  json:"category"`
	Author     uint          `json:"author" bson:"author"`
	Authors    []Author      `json:"authors" bson:"authors"`
	Categories []uint        `bson:"categories"  json:"categories"`
	Category   []Category    `bson:"category"  json:"category"`
	Tags       []uint        `bson:"tags"  json:"tags"`
	Tag        []Tag         `bson:"tag"  json:"tag"`
	ImageId    bson.ObjectId `json:"image_id" bson:"image_id"`
	// Image     []DataPostImage    `json:"image" bson:"image"`
	Views     string     `json:"views" bson:"views"`
	CreatedBy string     `json:"created_by" bson:"created_by"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedBy string     `json:"update_by" bson:"update_by"`
	UpdatedAt *time.Time `json:"update_at" bson:"update_at"`
	DeletedBy string     `json:"deleted_by" bson:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

type PostSurvey struct {
	Suka  uint `json:"suka" bson:"suka"`
	Lucu  uint `json:"lucu" bson:"lucu"`
	Kaget uint `json:"kaget" bson:"kaget"`
	Sedih uint `json:"sedih" bson:"sedih"`
	Marah uint `json:"marah" bson:"marah"`
}
