package jwt

import (
	"time"
	"uptime/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		Name: user.Name,
		Family: user.Family,
		HasVerified: user.HasVerified(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}
