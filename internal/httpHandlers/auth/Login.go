package auth

import (
	"errors"
	"net/http"
	"uptime/internal/jwt"
	"uptime/internal/models/User"
	"uptime/internal/validations/auth"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	params := auth.LoginValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := User.User{}
	if err := user.Find("username = ? OR email = ?", params.Username, params.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect username or password"})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	// check if email/username and password are not correct, return error: "incorrect username or password"
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect username or password"})
		return
	}

	// else, generate a jwt and return it.
	token, err := jwt.Generate(user)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
