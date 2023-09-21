package Otp

import "gorm.io/gorm"

type Otp struct {
	gorm.Model
	Email string `json:"email" gorm:"not null"`
	Code  int    `json:"code" gorm:"not null; unique"`
	Used  bool   `json:"used"`
}
