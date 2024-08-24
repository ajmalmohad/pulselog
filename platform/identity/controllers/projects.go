package controllers

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectRepository       *repositories.ProjectRepository
	projectMemberRepository *repositories.ProjectMemberRepository
}

func NewProjectController(
	projectRepository *repositories.ProjectRepository,
	projectMemberRepository *repositories.ProjectMemberRepository,
) *ProjectController {
	return &ProjectController{
		projectRepository:       projectRepository,
		projectMemberRepository: projectMemberRepository,
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

	// TODO: Make the creation of project and project member a transaction
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

	projectMember := &models.ProjectMember{
		ProjectID: project.ID,
		UserID:    userID,
		Role:      models.ADMIN,
	}

	if _, err := c.projectMemberRepository.Create(projectMember); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to create project member",
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

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project retrieved successfully",
		Data:    project,
	})
}

func (c *ProjectController) GetAllProjects(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	projects, err := c.projectRepository.FindAllByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to retrieve projects",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Projects retrieved successfully",
		Data:    projects,
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
