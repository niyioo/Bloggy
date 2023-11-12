package utils

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"my-blog-api/auth"
	"net/http"
	"regexp"
	"strings"
	"time"
)

<<<<<<< HEAD
=======
// RespondWithError sends an error response to the client
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

<<<<<<< HEAD
=======
	// Create an error response JSON
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

<<<<<<< HEAD
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
=======
	// Serialize and send the error response
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		// Handle serialization error
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
		log.Printf("Error encoding error response: %v", err)
	}
}

<<<<<<< HEAD
func HandleError(err error) {
	log.Printf("Error: %v", err)
}

=======
// HandleError handles errors by logging or reporting them
func HandleError(err error) {
	// Log the error
	log.Printf("Error: %v", err)

	// You can add additional error handling or reporting logic here,
	// such as sending error notifications to a monitoring system.
}

// DecodeJSONBody decodes the JSON request body
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func DecodeJSONBody(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(target); err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
=======
// EncodeJSONResponse encodes and sends a JSON response
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func EncodeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
=======
// SetJSONResponseHeaders sets response headers for JSON content
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func SetJSONResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

<<<<<<< HEAD
func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	SetJSONResponseHeaders(w)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
=======
// RespondWithJSON formats and sends JSON responses
func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	SetJSONResponseHeaders(w)

	// Serialize and send the JSON response
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// Handle serialization error
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
		RespondWithError(w, http.StatusInternalServerError, "Failed to encode JSON response")
	}
}

<<<<<<< HEAD
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

=======
// AuthMiddleware is a middleware function for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request headers or parameters
		tokenString := r.Header.Get("Authorization")

		// Check if the user is authenticated based on the JWT token
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
		isAuthenticated, err := auth.IsAuthenticated(tokenString)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if isAuthenticated {
<<<<<<< HEAD
=======
			// If authenticated, call the next handler
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
			next.ServeHTTP(w, r)
		} else {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		}
	})
}

<<<<<<< HEAD
func AuthorizationMiddleware(roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

=======
// AuthorizationMiddleware is a middleware function for role-based authorization
func AuthorizationMiddleware(roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request headers or parameters
		tokenString := r.Header.Get("Authorization")

		// Check if the user has one of the required roles based on the JWT token
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
		hasRoles, err := auth.HasRoles(tokenString, roles)
		if err != nil {
			RespondWithError(w, http.StatusForbidden, "Forbidden")
			return
		}

		if hasRoles {
<<<<<<< HEAD
=======
			// If authorized, call the next handler
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
			next.ServeHTTP(w, r)
		} else {
			RespondWithError(w, http.StatusForbidden, "Forbidden")
		}
	})
}

<<<<<<< HEAD
=======
// LogError logs errors along with additional context or messages
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func LogError(message string, err error) {
	log.Printf("Error: %s - %v", message, err)
}

<<<<<<< HEAD
=======
// LogInfo logs informational messages or events
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func LogInfo(message string) {
	log.Printf("Info: %s", message)
}

<<<<<<< HEAD
=======
// LogDebug logs debug information for debugging purposes
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func LogDebug(message string) {
	log.Printf("Debug: %s", message)
}

<<<<<<< HEAD
=======
// FormatDateTime formats a time.Time object as a string
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func FormatDateTime(t time.Time, format string) string {
	return t.Format(format)
}

<<<<<<< HEAD
func IsValidEmail(email string) bool {
=======
// IsValidEmail checks if an email address is valid
func IsValidEmail(email string) bool {
	// A basic email validation using regular expression
	// You can use a more robust regular expression for validation
	// This is a simple example, not foolproof
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
	valid := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	return valid.MatchString(email)
}

<<<<<<< HEAD
=======
// GenerateRandomString generates a random string of a specified length
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < length; i++ {
		index := r.Intn(len(charset))
		result.WriteByte(charset[index])
	}

	return result.String()
}

<<<<<<< HEAD
func ValidateInput(input string) error {
=======
// ValidateInput validates input data according to specific criteria
func ValidateInput(input string) error {
	// Check if the input is at least 5 characters long
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
	if len(input) < 5 {
		return errors.New("Input must be at least 5 characters long")
	}

<<<<<<< HEAD
=======
	// You can add more validation rules here as needed
	// For example, check for specific characters, patterns, or other criteria

>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
	return nil
}
