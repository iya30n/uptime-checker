package models

import (
	"time"

	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null; size:100"`
	Url       string    `json:"url" gorm:"not null"`
	CheckTime time.Duration `json:"check_time" gorm:"not null"`
	UserId uint `json:"user_id" gorm:"not null;"`
	User   User	`json:"-"`

	preloads []string `json:"-" gorm:"-"`
}

func (w *Website) With(preloads []string) {
	w.preloads = preloads
}

func (w Website) All() ([]Website, error) {
	var websites []Website
	mydb := db
	for _, p := range w.preloads {
		mydb = db.Preload(p)
	}

	res := mydb.Find(&websites)

	return websites, res.Error
}

func (w Website) Get(userId uint) ([]Website, error) {
	var websites []Website

	mydb := db
	for _, p := range w.preloads {
		mydb = db.Preload(p)
	}

	res := mydb.Where("user_id = ?", userId).Find(&websites)

	return websites, res.Error
}

func (w *Website) First(query string, values ...interface{}) error {
	mydb := db
	for _, p := range w.preloads {
		mydb = db.Preload(p)
	}

	return mydb.Where(query, values...).First(&w).Error
}

func (w *Website) Store() error {
	return db.Create(&w).Error
}

func (w *Website) Update(data map[string]interface{}) error {
	return db.Model(&w).Updates(data).Error
}

func (w *Website) Delete() error {
	return db.Delete(&w).Error
}