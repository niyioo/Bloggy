package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("12345")

type Claims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

func GenerateToken(userID, userRole string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func IsAuthenticated(tokenString string) (bool, error) {
	_, err := ParseToken(tokenString)
	return err == nil, err
}

func HasRoles(tokenString string, roles []string) (bool, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role == "admin" && claims.UserRole == "admin" {
			return true, nil
		} else if role == "user" && claims.UserRole == "user" {
			return true, nil
		}
	}

	return false, nil
}
