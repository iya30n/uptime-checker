package auth

import (
	"fmt"
	"net/http"
	"uptime/internal/validations/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	params := auth.LoginValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err.Error())})
		return
	}

	c.JSON(200, params)
	// check if email/username and password are not correct, return error: "invalid username or password"
	// else, generate a jwt and return it.
}