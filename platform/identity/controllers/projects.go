package controllers

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectRepository *repositories.ProjectRepository
	userRepository    *repositories.UserRepository
}

func NewProjectController(
	projectRepository *repositories.ProjectRepository,
	userRepository *repositories.UserRepository,
) *ProjectController {
	return &ProjectController{
		projectRepository: projectRepository,
		userRepository:    userRepository,
	}
}

func (c *ProjectController) CreateProject(ctx *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
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
		return
	}

	user, err := c.userRepository.FindByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "User not found",
			Detail: err.Error(),
		})
		return
	}

	project := &models.Project{
		Name:    input.Name,
		OwnerID: user.ID,
	}

	if _, err := c.projectRepository.Create(project); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to create project",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project created successfully",
		Data:    project,
	})
}

func (c *ProjectController) GetProject(ctx *gin.Context) {
	projectID := ctx.Query("project_id")
	if projectID == "" {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: "Project ID is required",
		})
		return
	}

	pid, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid project ID",
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

	// Only project owner or project member can view the project
	project, err := c.projectRepository.FindByIDAndUser(uint(pid), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project not found",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project retrieved successfully",
		Data:    project,
	})
}

func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	// update project
}

func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	// delete project
}
