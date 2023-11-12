package routes

import (
	"context"
	"errors"
	"my-blog-api/database"
	"my-blog-api/models"
	"my-blog-api/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBlogPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/blogposts", CreateBlogPostHandler).Methods("POST")
	router.HandleFunc("/api/blogposts", GetBlogPosts).Methods("GET")
	router.HandleFunc("/api/blogposts/{id}", GetBlogPost).Methods("GET")
}

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

	collection := client.Database("bloggy").Collection("blog_posts")

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

func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := getPostByID(objectID)
	if err != nil {
		if err == ErrPostNotFound {
			utils.RespondWithError(w, http.StatusNotFound, "Blog post not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get blog post")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, post)
}

var ErrPostNotFound = errors.New("blog post not found")

func getPostByID(objectID primitive.ObjectID) (models.BlogPost, error) {
	client, err := database.Connect()
	if err != nil {
		return models.BlogPost{}, err
	}
	defer database.Close(client)

	collection := client.Database("bloggy").Collection("blog_posts")

	var post models.BlogPost
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.BlogPost{}, ErrPostNotFound
		}
		return models.BlogPost{}, err
	}

	return post, nil
}
