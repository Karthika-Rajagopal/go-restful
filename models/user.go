package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Email           string         `gorm:"unique;not null" json:"email"`
	Password        string         `gorm:"not null" json:"-"`
	MetamaskAddress string         `gorm:"not null" json:"metamask_address"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id string) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func CreateUser(user *User) error {
	return db.Create(user).Error
}

// UpdateUserMetamaskAddress updates the user's metamask address
func UpdateUserMetamaskAddress(userId, metamaskAddress string) error {
	return db.Model(&User{}).Where("id = ?", userId).Update("metamask_address", metamaskAddress).Error
}
