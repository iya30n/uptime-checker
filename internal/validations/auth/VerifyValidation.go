package auth

import "uptime/internal/validations"

type VerifyValidation struct {
	validations.Parser

	Email string `json:"email" binding:"required,email,min=8,max=70"`
	Code  int    `json:"code" binding:"required,min=5"`
}
