package auth

import "uptime/internal/validations"

type ResendOtpValidation struct {
	validations.Parser

	Email string `json:"email" binding:"required,email,min=8,max=70"`
}
