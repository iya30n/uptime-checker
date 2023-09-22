package Otp

import (
	"errors"
	"math/rand"

	"gorm.io/gorm"
)

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