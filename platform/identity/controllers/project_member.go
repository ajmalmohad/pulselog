package controllers

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

	"github.com/gin-gonic/gin"
)

type ProjectMemberController struct {
	projectRepository       *repositories.ProjectRepository
	projectMemberRepository *repositories.ProjectMemberRepository
}

func NewProjectMemberController(
	projectRepository *repositories.ProjectRepository,
	projectMemberRepository *repositories.ProjectMemberRepository,
) *ProjectMemberController {
	return &ProjectMemberController{
		projectRepository:       projectRepository,
		projectMemberRepository: projectMemberRepository,
	}
}

func (c *ProjectMemberController) CreateProjectMember(ctx *gin.Context) {
	var input struct {
		ProjectID uint        `json:"project_id" binding:"required"`
		UserID    uint        `json:"user_id" binding:"required"`
		Role      models.Role `json:"role" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request body",
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
		ctx.Abort()
		return
	}

	isOwner, err := c.projectRepository.IsOwner(input.ProjectID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to check if user is owner of project",
			Detail: err.Error(),
		})
		return
	}

	if !isOwner {
		ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{
			Error:  "You are not authorized to create project members",
			Detail: "Owner access required",
		})
		return
	}

	projectMember := &models.ProjectMember{
		ProjectID: input.ProjectID,
		UserID:    input.UserID,
		Role:      input.Role,
	}

	if _, err := c.projectMemberRepository.Create(projectMember); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to create project member",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project member created successfully",
		Data:    projectMember,
	})
}

func (c *ProjectMemberController) GetProjectMember(ctx *gin.Context) {
	projectMemberID, err := utils.GetProjectMemberIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	projectMember, err := c.projectMemberRepository.FindByID(projectMemberID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project member not found",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project member retrieved successfully",
		Data:    projectMember,
	})
}

func (c *ProjectMemberController) GetAllProjectMembers(ctx *gin.Context) {
	projectID, err := utils.GetProjectIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	projectMembers, err := c.projectMemberRepository.FindAllByProjectID(projectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to retrieve project members",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project members retrieved successfully",
		Data:    projectMembers,
	})
}

func (c *ProjectMemberController) UpdateProjectMember(ctx *gin.Context) {
	var input struct {
		Role models.Role `json:"role"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request body",
			Detail: err.Error(),
		})
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

	projectMember, err := c.projectMemberRepository.FindByID(projectMemberID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project member not found",
			Detail: err.Error(),
		})
		return
	}

	if input.Role != "" {
		projectMember.Role = input.Role
	}

	if _, err := c.projectMemberRepository.Update(projectMember); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to update project member",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project member updated successfully",
		Data:    projectMember,
	})
}

func (c *ProjectMemberController) DeleteProjectMember(ctx *gin.Context) {
	projectMemberID, err := utils.GetProjectMemberIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	projectMember, err := c.projectMemberRepository.FindByID(projectMemberID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project member not found",
			Detail: err.Error(),
		})
		return
	}

	if _, err := c.projectMemberRepository.Delete(projectMember); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete project member",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project member deleted successfully",
		Data:    projectMember,
	})
}
