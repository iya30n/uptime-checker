package website

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
	"uptime/internal/jwt"
	"uptime/internal/models"
	"uptime/pkg/influxdb"
	"uptime/pkg/logger"
)

func UptimeHistory(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	claims, err := jwt.Parse(token)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	websiteId := c.Param("id")
	website := models.Website{}
	if err := website.First("id = ? and user_id = ?", websiteId, claims.UserId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Website not found"})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}

	from, _ := time.Parse(time.DateTime, c.Query("from"))
	to, _ := time.Parse(time.DateTime, c.Query("to"))

	filter := influxdb.ReadFilters{
		UserId:    claims.UserId,
		WebsiteId: website.ID,
		FromDate:  from.Format(time.RFC3339),
		ToDate:    to.Format(time.RFC3339),
	}

	err, data := influxdb.Read(filter)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
