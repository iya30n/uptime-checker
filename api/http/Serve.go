package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"uptime/internal/httpHandlers/auth"
	"uptime/internal/middlewares"
	"net/http"
)

func Serve() {
	router := gin.Default()

	router.Use(cors.Default())
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	router.POST("/auth/register", auth.Register)
	router.POST("/auth/verify", auth.Verify)
	router.POST("/auth/resend-otp", middlewares.RateLimit(time.Minute*3, 1), auth.ResendOtp)
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/refresh-token", middlewares.Auth(), auth.RefreshToken)

	router.Run(":7000")
}
