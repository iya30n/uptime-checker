package http

import (
	"uptime/internal/httpHandlers/auth"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	router.POST("/auth/register", auth.Register)
	router.POST("/auth/verify", auth.Verify)
	router.POST("/auth/login", auth.Login)
		
	router.Run("localhost:7000")
}
