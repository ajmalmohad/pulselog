package controllers

import (
	"net/http"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository         *repositories.UserRepository
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewUserController(userRepository *repositories.UserRepository, refreshTokenRepository *repositories.RefreshTokenRepository) *UserController {
	return &UserController{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (u *UserController) DeleteUserHandler(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	user, err := u.userRepository.FindByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "User not found",
			Detail: err.Error(),
		})
		return
	}

	_, err = u.userRepository.Delete(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete user",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "User deleted successfully",
		Data:    user,
	})
}

func (u *UserController) LogoutUserHandler(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid input",
			Detail: err.Error(),
		})
		return
	}

	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	err = u.refreshTokenRepository.DeleteByTokenAndUserID(input.RefreshToken, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete refresh token",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "User logged out successfully",
	})
}

func (u *UserController) LogoutAllUserHandler(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	err = u.refreshTokenRepository.DeleteByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete all refresh tokens",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "User logged out from all devices successfully",
	})
}
