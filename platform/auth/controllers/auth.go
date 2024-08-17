package controllers

import (
	"pulselog/auth/repositories"

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
}

func (c *AuthController) LoginHandler(ctx *gin.Context) {
}

func (c *AuthController) ReauthenticateHandler(ctx *gin.Context) {
}

func (c *AuthController) DisableUserHandler(ctx *gin.Context) {
}

func (c *AuthController) DeleteUserHandler(ctx *gin.Context) {
}
