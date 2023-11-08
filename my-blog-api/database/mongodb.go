package database

import (
	"context"
	"log"
	"my-blog-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURL      = "mongodb://localhost:27017"
	clientOptions = options.Client().ApplyURI(mongoURL)
)

func Connect() (*mongo.Client, error) {
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create a new MongoDB client: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	return client, nil
}

func Close(client *mongo.Client) {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Printf("Failed to close MongoDB client: %v", err)
		}
	}
}

func IsEmailUnique(client *mongo.Client, email string) (bool, error) {
	collection := client.Database("bloggy").Collection("users")

	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func CreateUser(client *mongo.Client, email, hashedPassword string) (string, error) {
	collection := client.Database("bloggy").Collection("users")

	user := models.User{
		Email:          email,
		HashedPassword: hashedPassword,
		// Add other user properties as needed
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}

func GetUserByEmail(client *mongo.Client, email string) (*models.User, error) {
	collection := client.Database("bloggy").Collection("users")

	filter := bson.M{"email": email}
	var user models.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
