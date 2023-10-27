package http

import (
	"time"
	"uptime/internal/middlewares"

	"github.com/gin-gonic/gin"

	auth_handler "uptime/internal/httphandlers/auth"
	website_handler "uptime/internal/httphandlers/website"
)

func Serve() {
	router := gin.Default()

	// TODO: check these headers values for production
	router.Use(middlewares.CORS())

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
	website.PUT("/:id", website_handler.Update)
	website.DELETE("/:id", website_handler.Delete)

	router.Run(":7000")
}
