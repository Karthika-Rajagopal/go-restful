package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Karthika-Rajagopal/go-restful/models"
	"github.com/Karthika-Rajagopal/go-restful/utils"
)


// RegisterUser handles the registration of a new user
func RegisterUser(c *gin.Context) {
	var registerData models.RegisterRequest
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request data
	if err := validate.Struct(registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email already exists
	existingUser, err := models.GetUserByEmail(registerData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Create a new user
	user := &models.User{
		Email:    registerData.Email,
		Password: utils.HashPassword(registerData.Password),
	}

	if err := models.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// UserLogin handles the user login process
func UserLogin(c *gin.Context) {
	var loginData models.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request data
	if err := validate.Struct(loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	user, err := models.GetUserByEmail(loginData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the password
	if !utils.CheckPassword(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}