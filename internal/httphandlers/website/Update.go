package website

import (
	"errors"
	"net/http"
	"uptime/internal/validations/website"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context) {
	userId, err := getAuthId(c)
	if err != nil {
		return
	}

	params := website.UpdateValidation{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": params.Parse(err)})
		return
	}

	websiteId := c.Param("id")
	website, ok := getWebsite(c, userId, websiteId)
	if !ok {
		return
	}

	updateData := map[string]interface{}{
		"name":       params.Name,
		"url":        params.Url,
		"check_time": params.GetChcekTimeDur(),
		"notify":     *params.Notify,
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
