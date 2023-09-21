package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"uptime/internal/models/Otp"
	"uptime/internal/models/User"
	authvalidation "uptime/internal/validations/auth"
	"uptime/pkg/config"
	"uptime/pkg/logger"
	"uptime/pkg/mail"
	"uptime/pkg/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// TODO: use a queue for sending mails with retry
	if err := sendVerificationEmail(user.Email); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(201, gin.H{"message": "Thanks for your registration. Please check your email and verify your account"})
}

func sendVerificationEmail(email string) error {
	code, err := generateCode(email)
	if err != nil {
		return err
	}

	view := view.View{
		Path: "./views/mail/verify.html",
		Data: map[string]string{
			"[APP_URL]":           config.Get("APP_URL"),
			"[USER_EMAIL]":        email,
			"[VERIFICATION_CODE]": fmt.Sprintf("%d", code),
		},
	}

	return mail.Send(email, "Verification Email", view.Render())
}

func generateCode(email string) (int, error) {
	// TODO: check if code is unique in otp table
	vc := Otp.Otp{
		Email: email,
		Code:  rand.Intn(99999-10123) + 10123,
	}

	err := vc.Save()
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return generateCode(email)
	}

	return vc.Code, err
}
