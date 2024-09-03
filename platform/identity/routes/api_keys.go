package routes

import (
	"pulselog/identity/controllers"
	"pulselog/identity/middleware"
	"pulselog/identity/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAPIKeysRoutes(router *gin.Engine, db *gorm.DB) {
	apiKeysRouter := router.Group("/api-keys")
	userRepository := repositories.NewUserRepository(db)
	apiKeyRepository := repositories.NewAPIKeyRepository(db)
	projectRepository := repositories.NewProjectRepository(db)
	apiKeysRouter.Use(middleware.AuthMiddleware(userRepository))
	apiKeyController := controllers.NewAPIKeyController(apiKeyRepository)

	apiKeysRouter.POST("", middleware.ProjectMemberOnly(projectRepository), apiKeyController.CreateAPIKey)
}
