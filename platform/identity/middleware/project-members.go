package middleware

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

	"github.com/gin-gonic/gin"
)

func ProjectMemberMiddleware(
	projectMemberRepository *repositories.ProjectMemberRepository,
	adminOnly bool,
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

		projectMemberID, err := utils.GetProjectMemberIDFromQuery(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
				Error:  "Invalid request",
				Detail: err.Error(),
			})
			return
		}

		projectMember, err := projectMemberRepository.FindByID(projectMemberID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, types.ErrorResponse{
				Error:  "Failed to get project member",
				Detail: err.Error(),
			})
			return
		}

		allMembers, err := projectMemberRepository.FindAllByProjectID(projectMember.ProjectID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, types.ErrorResponse{
				Error:  "Failed to get project members",
				Detail: err.Error(),
			})
			return
		}

		for _, member := range allMembers {
			if member.UserID == userID {
				if adminOnly && !(member.Role == models.ADMIN) {
					ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{
						Error:  "You are not authorized to access this project member",
						Detail: "Admin access required",
					})
					ctx.Abort()
					return
				}
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{
			Error:  "You are not authorized to access this project member",
			Detail: "You are not a member of this project",
		})
	}
}

func SameProjectAdminOnly(
	projectMemberRepository *repositories.ProjectMemberRepository,
) gin.HandlerFunc {
	return ProjectMemberMiddleware(projectMemberRepository, true)
}

func SameProjectMemberOnly(
	projectMemberRepository *repositories.ProjectMemberRepository,
) gin.HandlerFunc {
	return ProjectMemberMiddleware(projectMemberRepository, false)
}
