package jwt

import "github.com/golang-jwt/jwt/v4"

func Verify(token string) bool {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return false
	}

	return tkn.Valid
}
