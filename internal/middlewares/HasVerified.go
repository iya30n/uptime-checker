package middlewares

import (
	"net/http"
	"strings"
	"uptime/internal/jwt"
	"uptime/internal/models"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
)

func HasVerified() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

		claims, err := jwt.Parse(token)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
			return
		}

		user := models.User{}
		if err := user.First("id = ?", claims.UserId); err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
			return
		}

		if !user.HasVerified() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please verify your account first."})
			return
		}

		c.Next()
	}
}
