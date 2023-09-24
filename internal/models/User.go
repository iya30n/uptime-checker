package models

import (
	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name            string    `json:"name" gorm:"not null"`
	Family          string    `json:"family" gorm:"not null"`
	Email           string    `json:"email" gorm:"not null; unique"`
	Username        string    `json:"username" gorm:"not null; unique"`
	Password        string    `json:"password" gorm:"not null"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Websites        []Website
}

func (u *User) First(query string, values ...interface{}) error {
	return db.Where(query, values...).First(&u).Error
}

func (u *User) Save() error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(passwd)

	return db.Create(&u).Error
}

func (u *User) Update(data map[string]interface{}) error {
	return db.Model(&u).Updates(data).Error
}

func (u *User) Exists() (bool, error) {
	res := db.Where("username = ? OR email = ?", u.Username, u.Email).Find(&User{})

	return res.RowsAffected != 0, res.Error
}

func (u *User) HasVerified() bool {
	return !u.EmailVerifiedAt.IsZero()
}
