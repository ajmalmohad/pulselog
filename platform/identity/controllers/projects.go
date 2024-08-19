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
}

func NewProjectController(
	projectRepository *repositories.ProjectRepository,
) *ProjectController {
	return &ProjectController{
		projectRepository: projectRepository,
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

	project := &models.Project{
		Name:    input.Name,
		OwnerID: userID,
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

	project, err := c.projectRepository.FindByID(uint(pid))
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
	var input struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request body",
			Detail: err.Error(),
		})
		return
	}

	projectID, err := utils.GetProjectIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	project, err := c.projectRepository.FindByID(projectID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project not found",
			Detail: err.Error(),
		})
		return
	}

	if input.Name != "" {
		project.Name = input.Name
	}

	if _, err := c.projectRepository.Update(project); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to update project",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project updated successfully",
		Data:    project,
	})
}

func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	projectID, err := utils.GetProjectIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	project, err := c.projectRepository.FindByID(projectID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "Project not found",
			Detail: err.Error(),
		})
		return
	}

	if _, err := c.projectRepository.Delete(project); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete project",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project deleted successfully",
		Data:    project,
	})
}
