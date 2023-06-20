package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful/models"
	"github.com/Karthika-Rajagopal/go-restful/config"
	"github.com/Karthika-Rajagopal/go-restful/utils"
)

// LoginRequest represents the request body for the login API
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response body for the login API
type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// Login handles the login API
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := utils.VerifyPassword(req.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := LoginResponse{
		User:  user,
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}

//UpdateProfileRequest represents the request body for the update profile API
type UpdateProfileRequest struct {
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// UpdateProfile updates the user's profile.
// It extracts the public address from the signed message and updates it in the database.
// Only the metamask address is updated, and no address is updated from the request.
// @Summary Update User Profile
// @Description Update User Profile API
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body UpdateProfileRequest true "Update Profile Request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /user/profile [put]

func UpdateProfile(c *gin.Context) {
	// Get the current user from the JWT token
	user, _ := c.Get("user")

	// Extract the public address from the signed message and update it in the database
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicAddress, err := models.ExtractPublicAddress(req.Message, req.Signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to extract public address"})
		return
	}

	user.(*models.User).MetamaskAddress = publicAddress

	// Save the updated user in the database
	db := config.GetDB()
	db.Save(user.(*models.User))

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}