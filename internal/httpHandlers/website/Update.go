package website

import (
	"errors"
	"net/http"
	"uptime/internal/models"
	"uptime/internal/validations/website"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context) {
	params := website.UpdateValidation{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	websiteId := c.Param("id")
	website := models.Website{}
	if err := website.First("id = ?", websiteId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Website not found"})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Website not found"})
		return
	}

	updateData := map[string]interface{}{
		"name":       params.Name,
		"url":        params.Url,
		"check_time": params.CheckTime,
	}

	if err := website.Update(updateData); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"message": "Website already exists"})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Website updated successfully", "data": website})
}
