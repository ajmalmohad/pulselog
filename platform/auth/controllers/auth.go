package controllers

import (
	"net/http"
	"pulselog/auth/models"
	"pulselog/auth/repositories"
	"pulselog/auth/types"
	"pulselog/auth/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userRepository         *repositories.UserRepository
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewAuthController(userRepository *repositories.UserRepository, refreshTokenRepository *repositories.RefreshTokenRepository) *AuthController {
	return &AuthController{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (c *AuthController) SignupHandler(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to hash password"})
		return
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	_, err = c.userRepository.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create user"})
		return
	}

	accessToken, err := utils.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create access token"})
		return
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create refresh token"})
		return
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	_, err = c.refreshTokenRepository.Create(refreshTokenModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to save refresh token"})
		return
	}

	ctx.JSON(http.StatusCreated, types.SuccessResponse{
		Message: "User created successfully",
		Data: types.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func (c *AuthController) LoginHandler(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := c.userRepository.FindByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	accessToken, err := utils.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create access token"})
		return
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create refresh token"})
		return
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	_, err = c.refreshTokenRepository.Create(refreshTokenModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to save refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Login successful",
		Data: types.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func (c *AuthController) ReauthenticateHandler(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	err := utils.VerifyToken(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Invalid refresh token"})
		return
	}

	_, err = c.refreshTokenRepository.FindByToken(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Invalid refresh token"})
		return
	}

	claims, err := utils.ExtractClaims(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract claims"})
		ctx.Abort()
		return
	}

	userID, ok := claims["user_id"].(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract user ID from claims"})
		ctx.Abort()
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract email from claims"})
		ctx.Abort()
		return
	}

	accessToken, err := utils.CreateAccessToken(userID, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to create access token"})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Reauthentication successful",
		Data: types.TokenResponse{
			AccessToken: accessToken,
		},
	})
}
