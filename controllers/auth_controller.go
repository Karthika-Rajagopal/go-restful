package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful/models"
	"github.com/Karthika-Rajagopal/go-restful/config"
	"github.com/Karthika-Rajagopal/go-restful/utils"
)

// RegisterRequest represents the request body for the register API
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterResponse represents the response body for the register API
type RegisterResponse struct {
	PublicAddress string `json:"public_address"`
}

// Register handles the register API
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	response := RegisterResponse{
		PublicAddress: user.MetamaskAddr,
	}

	c.JSON(http.StatusOK, response)
}
