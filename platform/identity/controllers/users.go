package controllers

import (
	"net/http"
	"pulselog/identity/repositories"
	"pulselog/identity/types"

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

func (u *UserController) DeleteUserHandler(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "User ID not provided",
			Detail: "The user_id parameter is missing from the context",
		})
		return
	}

	user, err := u.userRepository.FindByID(userID.(uint))
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
	})
}
