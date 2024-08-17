package utils

import (
	"fmt"
	"pulselog/auth/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId, email string, expiryDuration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(expiryDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateAccessToken(userId, email string) (string, error) {
	return CreateToken(userId, email, time.Hour*24*7)
}

func CreateRefreshToken(userId, email string) (string, error) {
	return CreateToken(userId, email, time.Hour*24*30)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
