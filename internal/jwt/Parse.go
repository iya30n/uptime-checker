package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

func Parse(token string) (Claims, error) {
    claims := Claims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return claims, err
	}

	return claims, nil
}