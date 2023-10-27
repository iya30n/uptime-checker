package models

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"uptime/pkg/redis"
)

type Otp struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

func (o *Otp) Get(email string, code int) error {
	var err error

	scmd := redis.Connect().Get(context.Background(), email)
	if err = scmd.Err(); err != nil {
		return err
	}

	o.Email = email
	o.Code, err = scmd.Int()

	if o.Code != code {
		err = errors.New("Invalid code")
	}

	return err
}

func (o *Otp) Set() error {
	res := redis.Connect().Set(context.Background(), o.Email, o.Code, 3*time.Minute)

	return res.Err()
}

func (Otp) GenerateCode(email string) (int, error) {
	oc := Otp{
		Email: email,
		Code:  rand.Intn(99999-10123) + 10123,
	}

	err := oc.Set()

	return oc.Code, err
}
