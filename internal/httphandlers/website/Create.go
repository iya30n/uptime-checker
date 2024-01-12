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

func Create(c *gin.Context) {
	params := website.CreateValidation{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": params.Parse(err)})
		return
	}

	userId, err := getAuthId(c)
	if err != nil {
		return
	}

	website := models.Website{
		Name:      params.Name,
		Url:       params.Url,
		CheckTime: params.GetChcekTimeDur(),
		UserId:    userId,
		Notify:    *params.Notify,
	}

	if err := website.Store(); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Website already exists"})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Website created successfully", "data": website})
}
