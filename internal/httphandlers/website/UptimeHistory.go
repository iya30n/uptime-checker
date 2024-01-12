package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"uptime/internal/models"
	"uptime/pkg/influxdb"
	"uptime/pkg/logger"
)

func UptimeHistory(c *gin.Context) {
	userId, err := getAuthId(c)
	if err != nil {
		return
	}

	websiteId := c.Param("id")
	website, ok := getWebsite(c, userId, websiteId)
	if !ok {
		return
	}

	err, data := influxdb.Read(getInfluxReadFilter(c, website))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func getInfluxReadFilter(c *gin.Context, website models.Website) influxdb.ReadFilters {
	from, _ := time.Parse(time.DateTime, c.Query("from"))
	to, _ := time.Parse(time.DateTime, c.Query("to"))

	return influxdb.ReadFilters{
		UserId:    website.UserId,
		WebsiteId: website.ID,
		FromDate:  from.Format(time.RFC3339),
		ToDate:    to.Format(time.RFC3339),
	}
}
