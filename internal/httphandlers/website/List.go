package website

import (
	"net/http"
	"strings"
	"uptime/internal/jwt"
	"uptime/internal/models"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	claims, err := jwt.Parse(token)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	websites, err := new(models.Website).Get(claims.UserId)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}
