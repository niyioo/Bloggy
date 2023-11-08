package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("12345")

// Claims represents the custom claims you want to include in the JWT.
type Claims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

// GenerateToken generates a new JWT token for the specified user ID and user role.
func GenerateToken(userID string, userRole string) (string, error) {
	// Create the claims for the JWT token
	claims := &Claims{
		UserID:   userID,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create and sign the token using the HMAC SHA256 method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses and verifies a JWT token and returns the claims, including the user role.
func ParseToken(tokenString string) (*Claims, error) {
	// Parse the token with the custom claims type
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// IsAuthenticated checks if a user is authenticated based on a valid JWT token.
func IsAuthenticated(tokenString string) (bool, error) {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// HasRoles checks if a user has the required roles based on JWT claims.
func HasRoles(tokenString string, roles []string) (bool, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}

	// Check if the user has one of the required roles
	for _, role := range roles {
		if role == "admin" && claims.UserRole == "admin" {
			return true, nil
		} else if role == "user" && claims.UserRole == "user" {
			return true, nil
		}
	}

	return false, nil
}
