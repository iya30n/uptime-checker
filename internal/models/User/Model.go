package User

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string    `json:"name" gorm:"not null"`
	Family          string    `json:"family" gorm:"not null"`
	Email           string    `json:"email" gorm:"not null; unique"`
	Username        string    `json:"username" gorm:"not null; unique"`
	Password        string    `json:"password" gorm:"not null"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	// Websites []Website.Website
}
