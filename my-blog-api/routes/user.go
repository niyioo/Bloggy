package routes

import (
	"log"
	"my-blog-api/auth"
	"my-blog-api/database"
	"my-blog-api/models"
	"my-blog-api/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router *mux.Router) {
	// Define routes for user registration and login
	router.HandleFunc("/api/register", RegisterUser).Methods("POST")
	router.HandleFunc("/api/login", LoginUser).Methods("POST")
}

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get user registration data
	var userInput models.UserInput
	err := utils.DecodeJSONBody(r, &userInput)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate and process user registration
	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	if !isEmailUnique(client, userInput.Email) {
		// Hash the password for security (you can use a library like golang.org/x/crypto/bcrypt)
		hashedPassword := hashPassword(userInput.Password)

		// Create a new user in the database with the email and hashed password
		userID, err := createUser(client, userInput.Email, hashedPassword)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}

		// Generate a JWT token for the registered user
		var jwtKey = "12345"
		token, err := auth.GenerateToken(userID, jwtKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		// Respond with the token
		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"token": token})
	} else {
		utils.RespondWithError(w, http.StatusConflict, "Email already exists")
	}
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get user login data
	var userInput models.UserInput
	err := utils.DecodeJSONBody(r, &userInput)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate and process user login
	client, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to the database")
		return
	}
	defer database.Close(client)

	userID, err := authenticateUser(client, userInput.Email, userInput.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Authentication failed")
		return
	}

	// Generate a JWT token for the logged-in user
	var jwtKey = "12345"
	token, err := auth.GenerateToken(userID, jwtKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Respond with the token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// isEmailUnique checks if the email is unique in the database.
func isEmailUnique(client *mongo.Client, email string) bool {
	exists, err := database.IsEmailUnique(client, email)
	if err != nil {
		log.Printf("Error checking email uniqueness: %v", err)
		return false
	}
	return !exists
}

// hashPassword securely hashes the provided password.
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return ""
	}
	return string(hashedPassword)
}

// createUser creates a new user in the database with the provided email and hashed password.
func createUser(client *mongo.Client, email, hashedPassword string) (string, error) {
	userID, err := database.CreateUser(client, email, hashedPassword)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return "", err
	}
	return userID, nil
}

// authenticateUser authenticates the user based on the provided email and password.
func authenticateUser(client *mongo.Client, email, password string) (string, error) {
	user, err := database.GetUserByEmail(client, email)
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return "", err
	}

	userID := user.ID.Hex()
	return userID, nil
}
