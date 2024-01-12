package website

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"uptime/internal/models"
	"uptime/pkg/logger"
)

func getWebsite(c *gin.Context, userId uint, websiteId string) (models.Website, bool) {
	website := models.Website{}
	if err := website.First("id = ? and user_id = ?", websiteId, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Website not found"})
			return website, false
		}

		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return website, false
	}

	return website, true
}
