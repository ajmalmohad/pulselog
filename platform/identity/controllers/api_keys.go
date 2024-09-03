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
	var input struct {
		ProjectID uint `json:"project_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
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

	generatedKey, err := utils.CreateAPIToken(userID, input.ProjectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to generate API key",
			Detail: err.Error(),
		})
		return
	}

	apiKey := &models.APIKey{
		ProjectID: input.ProjectID,
		CreatedBy: userID,
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

func (a *APIKeyController) GetAPIKeys(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	apiKeys, err := a.apiKeyRepository.GetAPIKeysByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to get API keys",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "API keys for the user retrieved successfully",
		Data:    apiKeys,
	})
}

func (a *APIKeyController) DeleteAPIKey(ctx *gin.Context) {
	userID, _, err := utils.ExtractClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	apiKeyID, err := utils.GetAPIKeyIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:  "Invalid request",
			Detail: err.Error(),
		})
		return
	}

	apiKey, err := a.apiKeyRepository.FindByID(apiKeyID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, types.ErrorResponse{
			Error:  "API key not found",
			Detail: err.Error(),
		})
		return
	}

	if apiKey.CreatedBy != userID {
		ctx.JSON(http.StatusForbidden, types.ErrorResponse{
			Error:  "Forbidden",
			Detail: "You are not allowed to delete this API key",
		})
		return
	}

	result, err := a.apiKeyRepository.Delete(apiKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to delete API key",
			Detail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.SuccessResponse{
		Message: "API key deleted successfully",
		Data:    result,
	})
}
