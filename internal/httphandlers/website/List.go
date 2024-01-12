package website

import (
	"net/http"
	"uptime/internal/models"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	userId, err := getAuthId(c)
	if err != nil {
		return
	}

	websites, err := new(models.Website).Get(userId)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}
