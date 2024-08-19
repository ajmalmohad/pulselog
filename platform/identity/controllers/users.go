package controllers

import (
	"net/http"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

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
