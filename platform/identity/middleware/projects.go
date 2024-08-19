package middleware

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectMiddleware(
	projectRepository *repositories.ProjectRepository,
	allowedRoles []models.Role,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _, err := utils.ExtractClaimsFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error:  "Failed to extract user ID and email from claims",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		projectID, err := strconv.ParseUint(ctx.Param("projectID"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
				Error:  "Invalid project ID",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		_, err = projectRepository.FindByIDUserAndRoles(uint(projectID), userID, allowedRoles)
		if err != nil {
			ctx.JSON(http.StatusNotFound, types.ErrorResponse{
				Error:  "Project not found or you are not authorized",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func ProjectAdminOnly(
	projectRepository *repositories.ProjectRepository,
) gin.HandlerFunc {
	return ProjectMiddleware(projectRepository, []models.Role{models.ADMIN})
}

func ProjectMemberOnly(
	projectRepository *repositories.ProjectRepository,
) gin.HandlerFunc {
	return ProjectMiddleware(projectRepository, []models.Role{models.ADMIN, models.MEMBER})
}
