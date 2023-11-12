package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	// Add additional fields as needed, such as tags, categories, etc.
}

type BlogPostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	// Add other input fields as necessary.
}

func (input BlogPostInput) Validate() bool {
	if len(input.Title) == 0 || len(input.Content) == 0 {
		return false
	}

	// Add more validation rules here if needed.

	return true
}
