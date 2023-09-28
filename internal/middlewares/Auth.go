package middlewares

import (
	"net/http"
	"strings"
	"uptime/internal/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

		if len(token) < 1 || !jwt.Verify(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please Login first."})
			return
		}

		c.Next()
	}
}
