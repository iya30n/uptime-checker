package auth

type ResendOtpValidation struct {
	Email string `json:"email" binding:"required,email,min=8,max=70"`
}