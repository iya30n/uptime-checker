package Website

import (
	"time"
	"uptime/internal/models/User"

	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	Name string `json:"name" gorm:"not null; size:100"`
	Url string `json:"url" gorm:"not null"`
	CheckTime time.Time `json:"check_time" gorm:"not null"`

	UserId uint `json:"user_id" gorm:"not null;"`
	User User.User
}
