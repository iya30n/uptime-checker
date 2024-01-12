package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"uptime/internal/jwt"
	"uptime/pkg/logger"
)

func getAuthId(c *gin.Context) (uint, error) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	claims, err := jwt.Parse(token)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
	}

	return claims.UserId, err
}
