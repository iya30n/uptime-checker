package website

import (
	"net/http"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, err := getAuthId(c)
	if err != nil {
		return
	}

	websiteId := c.Param("id")
	website, ok := getWebsite(c, userId, websiteId)
	if !ok {
		return
	}

	if err := website.Delete(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Website deleted successfully"})
}
