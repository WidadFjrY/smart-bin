package helper

import (
	"os"
	"smart-trash-bin/domain/web"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func ValidationToken(tokenStr string) (*web.JwtClaims, error) {
	godotenv.Load()

	token, err := jwt.ParseWithClaims(tokenStr, &web.JwtClaims{}, func(tokenStr *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(*web.JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
