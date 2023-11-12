package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashedPassword"`
	// Add other user properties as needed
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// You can include other input fields as necessary.
}
