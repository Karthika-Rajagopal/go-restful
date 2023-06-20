package middlewares

import (
	"net/http"
	//"strings"

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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := utils.ExtractTokenFromHeader(authHeader)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := utils.VerifyToken(tokenString, config.GetJWTSecret())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, ok := token.Claims.(*utils.JWTClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Set user ID in the context for future use
		c.Set("userId", claims.UserID)

		c.Next()
	}
}