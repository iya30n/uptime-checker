package http

import (
	"log"
	"uptime/internal/httpHandlers/auth"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	router.POST("/auth/register", auth.Register)
	router.POST("/auth/login", auth.Login)

	log.Fatalln(router.Run(":9000"))
}