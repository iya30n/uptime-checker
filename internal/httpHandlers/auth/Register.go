package auth

import (
	"fmt"
	"net/http"
	"uptime/internal/models/User"
	authvalidation "uptime/internal/validations/auth"
	"uptime/pkg/mail"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	params := authvalidation.RegisterValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err.Error())})
		return
	}

	user := User.User{
		Name:     params.Name,
		Family:   params.Family,
		Email:    params.Email,
		Username: params.Username,
		Password: params.Password,
	}

	userExists, err := user.Exists()
	if err != nil {
		// TODO: log the error here
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	if userExists {
		c.JSON(http.StatusConflict, gin.H{"message": "Sorry! the username or email already exists."})
		return
	}

	if err := user.Save(); err != nil {
		// TODO: log the error here
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	// TODO: send an email to verify the account (use a queue for sending mails with retry)
	if err := mail.Send(user.Email, "Verification Email", "salam"); err != nil {
		// TODO: log the error here
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(201, gin.H{"message": "Thanks for your registration. Please check your email and verify your account"})
}
