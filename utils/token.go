package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// VerifyToken verifies the validity of a JWT token
func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}

// ExtractTokenFromHeader extracts the token string from the Authorization header
func ExtractTokenFromHeader(authorization string) string {
	if len(authorization) > 7 && authorization[:7] == "Bearer " {
		return authorization[7:]
	}
	return ""
}
