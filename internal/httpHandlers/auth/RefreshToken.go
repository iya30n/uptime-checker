package auth

import (
	"errors"
	"net/http"
	"uptime/internal/jwt"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	newToken, err := jwt.Refresh(token)
	if err != nil {
		if errors.Is(err, &jwt.TokenIsValidError{}) {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
