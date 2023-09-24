package models

import (
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	Email string `json:"email" gorm:"not null"`
	Code  int    `json:"code" gorm:"not null; unique"`
	Used  bool   `json:"used"`
}

func (o *Otp) First(email string, code int) error {
	return db.Where("email = ? AND code = ?", email, code).First(&o).Error
}

func (o *Otp) Save() error {
	return db.Create(&o).Error
}

func (o *Otp) Update(data map[string]interface{}) error {
	return db.Model(&o).Updates(data).Error
}

func (Otp) GenerateCode(email string) (int, error) {
	oc := Otp{
		Email: email,
		Code: rand.Intn(99999-10123) + 10123,
	}

	err := oc.Save()
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return oc.GenerateCode(email)
	}

	return oc.Code, err
}

func (o Otp) IsValid() bool {
	return !o.Used && o.CreatedAt.After(time.Now().Add(-time.Minute * 3))
}