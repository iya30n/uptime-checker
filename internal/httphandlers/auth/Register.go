package auth

import (
	"fmt"
	"net/http"
	"path/filepath"
	"uptime/internal/jobs"
	"uptime/internal/models"
	authvalidation "uptime/internal/validations/auth"
	"uptime/pkg/config"
	"uptime/pkg/logger"
	"uptime/pkg/queue"
	"uptime/pkg/view"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	params := authvalidation.RegisterValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": params.Parse(err)})
		return
	}

	user := models.User{
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

	if err := user.Save(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	if err := sendVerificationEmail(user.Email); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(201, gin.H{"message": "Thanks for your registration. Please check your email and verify your account"})
}

func sendVerificationEmail(email string) error {
	otp := models.Otp{}
	code, err := otp.GenerateCode(email)
	if err != nil {
		return err
	}

	view := view.View{
		Path: filepath.Join("views", "mail", "verify.html"),
		Data: map[string]string{
			"[APP_URL]":           config.Get("APP_URL"),
			"[USER_EMAIL]":        email,
			"[VERIFICATION_CODE]": fmt.Sprintf("%d", code),
		},
	}

	job := jobs.Job{
		Payload: jobs.JobPayload{
			"email": email,
			"title": "Verification Email",
			"view": view.Render(),
		},
	}

	return queue.Enqueue("upq:email", job)
}
