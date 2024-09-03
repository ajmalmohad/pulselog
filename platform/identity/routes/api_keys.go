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
	apiKeysRouter.Use(middleware.AuthMiddleware(userRepository))
	apiKeyController := controllers.NewAPIKeyController(apiKeyRepository)

	apiKeysRouter.POST("", apiKeyController.CreateAPIKey)
	apiKeysRouter.GET("/all", apiKeyController.GetAPIKeys) // Gets all api keys for the user
}
