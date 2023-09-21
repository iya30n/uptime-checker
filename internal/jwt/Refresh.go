package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Refresh(token string) (string, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return "", &TokenIsValidError{}
	}

	expireTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = jwt.NewNumericDate(expireTime)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newToken.SignedString(JwtKey)
}
