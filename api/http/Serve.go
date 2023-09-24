package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"uptime/internal/middlewares"
	"net/http"

	auth_handler "uptime/internal/httpHandlers/auth"
	website_handler "uptime/internal/httpHandlers/website"
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

	// auth routes
	auth := router.Group("/auth")
	auth.POST("/register", auth_handler.Register)
	auth.POST("/verify", auth_handler.Verify)
	auth.POST("/resend-otp", middlewares.RateLimit(time.Minute*3, 1), auth_handler.ResendOtp)
	auth.POST("/login", auth_handler.Login)
	auth.POST("/refresh-token", middlewares.Auth(), auth_handler.RefreshToken)

	// website routes
	website := router.Group("/website", middlewares.Auth(), middlewares.HasVerified())
	website.GET("/", website_handler.List)
	website.POST("/", website_handler.Create)

	router.Run(":7000")
}
