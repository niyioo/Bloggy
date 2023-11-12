package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogPost represents a blog post in the system.
type BlogPost struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	// You can include additional fields as needed, such as tags, categories, etc.
}

// BlogPostInput represents the input data for creating or updating a blog post.
type BlogPostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	// You can include other input fields as necessary.
}

// Validate validates the blog post input.
func (input BlogPostInput) Validate() bool {
	// Check if the title and content are not empty
	if len(input.Title) == 0 || len(input.Content) == 0 {
		return false
	}

	// You can add more validation rules here, for example:
	// - Ensure the title is not too long
	// - Ensure the content meets specific criteria
	// - Validate other fields if necessary

	// Return true if all validation rules pass
	return true
}
