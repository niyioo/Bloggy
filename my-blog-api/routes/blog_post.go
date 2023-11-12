package routes

import (
	"context"
	"my-blog-api/database"
	"my-blog-api/models"
	"my-blog-api/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBlogPostHandler is a handler function to create a new blog post.
func CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the blog post data
	var postInput models.BlogPostInput
	if err := utils.DecodeJSONBody(r, &postInput); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the blog post input
	if !postInput.Validate() {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid blog post data")
		return
	}

	// Create a new BlogPost instance using the input data
	newPost := models.BlogPost{
		ID:        primitive.NewObjectID(),
		Title:     postInput.Title,
		Content:   postInput.Content,
		Author:    "AuthorName", // You can set the author as needed
		CreatedAt: time.Now(),   // Set the creation timestamp
	}

	// Get the MongoDB client from the database package
	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	collection := client.Database("bloggy").Collection("blog_posts")

	// Insert the new blog post document
	insertResult, err := collection.InsertOne(context.Background(), newPost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a new blog post")
		return
	}

	// Respond with the newly created blog post as a JSON response
	utils.RespondWithJSON(w, http.StatusCreated, insertResult)
}

// GetBlogPosts retrieves a list of blog posts from MongoDB
func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	// Get the MongoDB client from the database package
	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	// Get a reference to the collection containing the blog posts
	collection := client.Database("your-database-name").Collection("blog_posts")

	filter := primitive.M{} // An empty filter to fetch all posts

	var posts []models.BlogPost
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch blog posts")
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var post models.BlogPost
		err := cur.Decode(&post)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode blog posts")
			return
		}
		posts = append(posts, post)
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}
