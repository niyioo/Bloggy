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

func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		log.Printf("Error encoding error response: %v", err)
	}
}

func HandleError(err error) {
	log.Printf("Error: %v", err)
}

func DecodeJSONBody(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(target); err != nil {
		return err
	}
	return nil
}

func EncodeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}

func SetJSONResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	SetJSONResponseHeaders(w)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to encode JSON response")
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		isAuthenticated, err := auth.IsAuthenticated(tokenString)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if isAuthenticated {
			next.ServeHTTP(w, r)
		} else {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		}
	})
}

func AuthorizationMiddleware(roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		hasRoles, err := auth.HasRoles(tokenString, roles)
		if err != nil {
			RespondWithError(w, http.StatusForbidden, "Forbidden")
			return
		}

		if hasRoles {
			next.ServeHTTP(w, r)
		} else {
			RespondWithError(w, http.StatusForbidden, "Forbidden")
		}
	})
}

func LogError(message string, err error) {
	log.Printf("Error: %s - %v", message, err)
}

func LogInfo(message string) {
	log.Printf("Info: %s", message)
}

func LogDebug(message string) {
	log.Printf("Debug: %s", message)
}

func FormatDateTime(t time.Time, format string) string {
	return t.Format(format)
}

func IsValidEmail(email string) bool {
	valid := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	return valid.MatchString(email)
}

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

func ValidateInput(input string) error {
	if len(input) < 5 {
		return errors.New("Input must be at least 5 characters long")
	}

	return nil
}
