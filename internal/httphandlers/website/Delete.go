package website

import (
	"errors"
	"net/http"
	"strings"
	"uptime/internal/jwt"
	"uptime/internal/models"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	claims, err := jwt.Parse(token)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	website := models.Website{}
	if err := website.First("id = ? and user_id = ?", c.Param("id"), claims.UserId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Website not found"})
			return
		}

		logger.Error(err.Error())
		c.JSON(404, gin.H{"message": "Website not found"})
		return
	}

	if err := website.Delete(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Website deleted successfully"})
}
