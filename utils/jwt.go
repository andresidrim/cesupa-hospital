package utils

import (
	"time"

	"github.com/andresidrim/cesupa-hospital/env"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(env.SECRET_KEY)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint(claims["user_id"].(float64))
		return id, nil
	}

	return 0, err
}
