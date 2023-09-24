package models

import (
	"time"

	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null; size:100"`
	Url       string    `json:"url" gorm:"not null"`
	CheckTime time.Time `json:"check_time" gorm:"not null"`

	UserId uint `json:"user_id" gorm:"not null;"`
	User   User
}

func (Website) Get(userId uint) ([]Website, error) {
	var websites []Website
	res := db.Where("user_id = ?", userId).Find(&websites)

	return websites, res.Error
}