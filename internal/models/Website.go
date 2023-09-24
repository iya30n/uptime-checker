package models

import (
	"time"

	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null; size:100"`
	Url       string    `json:"url" gorm:"not null; unique"`
	CheckTime time.Duration `json:"check_time" gorm:"not null"`

	UserId uint `json:"user_id" gorm:"not null;"`
	User   User	`json:"-"`
}

func (Website) Get(userId uint) ([]Website, error) {
	var websites []Website
	res := db.Where("user_id = ?", userId).Find(&websites)

	return websites, res.Error
}

func (w *Website) First(query string, values ...interface{}) error {
	return db.Where(query, values...).First(&w).Error
}

func (w *Website) Store() error {
	return db.Create(&w).Error
}

func (w *Website) Update(data map[string]interface{}) error {
	return db.Model(&w).Updates(data).Error
}