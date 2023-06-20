package utils

import (
	//"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/Karthika-Rajagopal/go-restful/config"
)

// JwtCustomClaims represents the custom claims used in JWT
type JwtCustomClaims struct {
	UserID uint
	jwt.StandardClaims
}

// GenerateJWTToken generates a new JWT token for the specified user ID
func GenerateJWTToken(userID uint) (string, error) {
	claims := &JwtCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		log.Printf("Failed to generate JWT token: %v", err)
		return "", err
	}

	return tokenString, nil
}

// VerifyJWTToken verifies the validity of a JWT token
func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Printf("Failed to verify JWT token: %v", err)
		return nil, err
	}

	return token, nil
}
