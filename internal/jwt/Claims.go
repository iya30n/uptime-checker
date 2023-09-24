package jwt

import (
	"uptime/pkg/config"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey []byte = []byte(config.Get("JWT_KEY"))

type Claims struct {
	UserId uint `json:"user_id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	HasVerified bool `json:"has_verified"`

	jwt.RegisteredClaims
}
