package middlewares

import (
	"net/http"
	"uptime/internal/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) < 1 || !jwt.Verify(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please Login first."})
			return
		}

		c.Next()
	}
}