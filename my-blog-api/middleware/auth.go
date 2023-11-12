package middleware

import (
	"my-blog-api/auth"
	"my-blog-api/utils"
	"net/http"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token is missing")
			return
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		context.Set(r, "userID", claims.UserID)

		next.ServeHTTP(w, r)
	})
}
