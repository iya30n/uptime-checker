package jwt

import (
	"uptime/pkg/config"

	"github.com/golang-jwt/jwt/v4"
)

// var JwtKey []byte = []byte("simple-app-jwt-key")
var JwtKey []byte = []byte(config.Get("JWT_KEY"))

type Claims struct {
	Username string
	jwt.RegisteredClaims
}
