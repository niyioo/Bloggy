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

func CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	var postInput models.BlogPostInput
	if err := utils.DecodeJSONBody(r, &postInput); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if !postInput.Validate() {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid blog post data")
		return
	}

	newPost := models.BlogPost{
		ID:        primitive.NewObjectID(),
		Title:     postInput.Title,
		Content:   postInput.Content,
		Author:    "AuthorName",
		CreatedAt: time.Now(),
	}

	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	collection := client.Database("bloggy").Collection("blog_posts")

	insertResult, err := collection.InsertOne(context.Background(), newPost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a new blog post")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, insertResult)
}

func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	collection := client.Database("your-database-name").Collection("blog_posts")

	filter := primitive.M{}

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
