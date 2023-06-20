package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	Password      string
	PublicAddress string
}

type UpdateProfileRequest struct {
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type UpdateProfileResponse struct {
	PublicAddress string `json:"public_address"`
}
