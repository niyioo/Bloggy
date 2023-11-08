package middleware

import (
	"my-blog-api/auth"
	"my-blog-api/utils"
	"net/http"

	"github.com/gorilla/context"
)

// AuthMiddleware is a middleware for JWT-based authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the request headers
		tokenString := r.Header.Get("Authorization")

		// Check if the token is present
		if tokenString == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Token is missing")
			return
		}

		// Parse the token and validate it
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Invalid token")
			return
		}

		// Set the user ID from the token in the request context
		context.Set(r, "userID", claims.UserID)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
