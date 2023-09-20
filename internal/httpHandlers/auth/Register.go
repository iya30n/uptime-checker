package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"uptime/internal/models/User"
	authvalidation "uptime/internal/validations/auth"
	"uptime/pkg/config"
	"uptime/pkg/logger"
	"uptime/pkg/mail"
	"uptime/pkg/view"

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
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	if userExists {
		c.JSON(http.StatusConflict, gin.H{"message": "Sorry! the username or email already exists."})
		return
	}

	// TODO: think about using transaction here (if sending mail filed, rollback the query).
	if err := user.Save(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	// TODO: send an email to verify the account (use a queue for sending mails with retry)
	if err := sendVerificationEmail(user); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(201, gin.H{"message": "Thanks for your registration. Please check your email and verify your account"})
}

func sendVerificationEmail(user User.User) error {
	code := rand.Intn(99999 - 10123) + 10123
	
	view := view.View{
		Path: "./views/mail/verify.html",
		Data: map[string]string{
			"[APP_URL]": config.Get("APP_URL"),
			"[USER_EMAIL]": user.Email,
			"[VERIFICATION_CODE]": fmt.Sprintf("%d", code),
		},
	}

	return mail.Send(user.Email, "Verification Email", view.Render())
}
