package controllers

import (
	"net/http"
	"pulselog/identity/models"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"

	"github.com/gin-gonic/gin"
)

type APIKeyController struct {
	apiKeyRepository *repositories.APIKeyRepository
}

func NewAPIKeyController(apiKeyRepository *repositories.APIKeyRepository) *APIKeyController {
	return &APIKeyController{
		apiKeyRepository: apiKeyRepository,
	}
}

func (a *APIKeyController) CreateAPIKey(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
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

	generatedKey, err := utils.CreateAPIToken(userID, projectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to generate API key",
			Detail: err.Error(),
		})
		return
	}

	apiKey := &models.APIKey{
		ProjectID: projectID,
		CreatedBy: &userID,
		Key:       generatedKey,
	}

	if _, err := a.apiKeyRepository.Create(apiKey); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to create API key",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "Project created successfully",
		Data:    types.APIKeyResponse{APIKey: generatedKey},
	})
}
