package http

import (
	"time"
	"uptime/internal/httpHandlers/auth"
	"uptime/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	router.POST("/auth/register", auth.Register)
	router.POST("/auth/verify", auth.Verify)
	router.POST("/auth/resend-otp", middlewares.RateLimit(time.Minute * 3, 1), auth.ResendOtp)
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/refresh-token", middlewares.Auth(), auth.RefreshToken)
		
	router.Run("localhost:7000")
}
