package auth

import "uptime/internal/validations"

type LoginValidation struct {
	validations.Parser
	Email string `json:"email" binding:"required_without=Username,omitempty,email,min=8,max=70"`
	//Username string `json:"username" binding:"required_without=Email,omitempty,min=5,max=70"`
	Username string `json:"username" binding:"omitempty,min=5,max=70"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}
