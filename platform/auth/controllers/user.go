package controllers

import (
	"net/http"
	"pulselog/auth/repositories"
	"pulselog/auth/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository *repositories.UserRepository
}

func NewUserController(userRepository *repositories.UserRepository) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

func (u *UserController) DisableUserHandler(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid user ID, must be a number"})
		return
	}

	tokenUserID := ctx.GetUint("user_id")
	if tokenUserID != uint(userID) {
		ctx.JSON(http.StatusForbidden, types.ErrorResponse{Error: "You are not allowed to disable this user"})
		return
	}

	user, err := u.userRepository.FindByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{Error: "User not found"})
		return
	}

	user.IsActive = false
	_, err = u.userRepository.Update(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to disable user"})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "User disabled successfully",
	})
}

func (u *UserController) DeleteUserHandler(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	tokenUserID := ctx.GetUint("user_id")
	if tokenUserID != uint(userID) {
		ctx.JSON(http.StatusForbidden, types.ErrorResponse{Error: "You are not allowed to disable this user"})
		return
	}

	user, err := u.userRepository.FindByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{Error: "User not found"})
		return
	}

	_, err = u.userRepository.Delete(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "User deleted successfully",
	})
}
