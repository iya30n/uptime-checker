package auth

import (
	"fmt"
	"net/http"
	"uptime/internal/models/User"
	"uptime/internal/validations/auth"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	params := auth.RegisterValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err.Error())})
		return
	}

	user := User.User{
		Name:     params.Name,
		Family:   params.Family,
		Email:    params.Email,
		Username: params.Username,
		Password: params.Password,
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	// TODO: send an email to verify the account (use a queue for sending mails)
}
