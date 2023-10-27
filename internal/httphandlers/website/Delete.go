package website

import (
	"errors"
	"net/http"
	"uptime/internal/models"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context) {
	website := models.Website{}
	if err := website.First("id = ?", c.Param("id")); err != nil {
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