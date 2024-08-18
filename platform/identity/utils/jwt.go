package utils

import (
	"fmt"
	"pulselog/auth/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint, email string, expiryDuration time.Duration) (string, error) {
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

func CreateAccessToken(userId uint, email string) (string, error) {
	return CreateToken(userId, email, time.Hour)
}

func CreateRefreshToken(userId uint, email string) (string, error) {
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

func ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func ExtractUserIDAndEmailFromClaims(claims map[string]interface{}) (uint, string, error) {
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("failed to extract user ID from claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return 0, "", fmt.Errorf("failed to extract email from claims")
	}

	return uint(userIDFloat), email, nil
}
