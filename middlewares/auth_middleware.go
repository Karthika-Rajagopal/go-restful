package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/Karthika-Rajagopal/go-restful/models"
	//"github.com/Karthika-Rajagopal/go-restful/config"
	"github.com/Karthika-Rajagopal/go-restful/utils"
)

// AuthMiddleware validates the JWT token and sets the user ID in the context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := utils.VerifyJWTToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*utils.JwtCustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token claims"})
			c.Abort()
			return
		}

		// Set the user ID in the context
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
