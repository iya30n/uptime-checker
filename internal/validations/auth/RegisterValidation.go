package auth

type RegisterValidation struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Family   string `json:"family" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,min=8,max=70"`
	Username string `json:"username" binding:"required,min=5,max=70"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}
