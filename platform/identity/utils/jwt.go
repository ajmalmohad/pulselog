package utils

import (
	"fmt"
	"pulselog/identity/config"
	"pulselog/identity/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(userId uint, email string, expiryDuration time.Duration) (string, error) {
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
	return createJWT(userId, email, time.Hour)
}

func CreateRefreshToken(userId uint, email string) (string, error) {
	return createJWT(userId, email, time.Hour*24*30)
}

func CreateTokens(userId uint, email string) (types.TokenResponse, error) {
	accessToken, err := CreateAccessToken(userId, email)
	if err != nil {
		return types.TokenResponse{}, err
	}

	refreshToken, err := CreateRefreshToken(userId, email)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return types.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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

func ExtractUserIDAndEmailFromClaims(tokenString string) (uint, string, error) {
	claims, err := extractClaims(tokenString)
	if err != nil {
		return 0, "", err
	}

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

func extractClaims(tokenString string) (jwt.MapClaims, error) {
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

func CreateAPIToken(userId uint, projectId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"project_id": projectId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractUserIDAndProjectIDFromClaims(tokenString string) (uint, uint, error) {
	claims, err := extractClaims(tokenString)
	if err != nil {
		return 0, 0, err
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("failed to extract user ID from claims")
	}

	projectIDFloat, ok := claims["project_id"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("failed to extract project ID from claims")
	}

	return uint(userIDFloat), uint(projectIDFloat), nil
}
