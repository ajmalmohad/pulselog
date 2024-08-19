package routes

import (
	"pulselog/identity/controllers"
	"pulselog/identity/middleware"
	"pulselog/identity/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProjectRoutes(router *gin.Engine, db *gorm.DB) {
	projectRouter := router.Group("/projects")
	projectRouter.Use(middleware.AuthMiddleware())
	projectRepository := repositories.NewProjectRepository(db)
	projectController := controllers.NewProjectController(projectRepository)

	projectRouter.POST("", projectController.CreateProject)
	projectRouter.GET("", middleware.ProjectMemberOnly(projectRepository), projectController.GetProject)
	projectRouter.PUT("", middleware.ProjectAdminOnly(projectRepository), projectController.UpdateProject)
	projectRouter.DELETE("", middleware.ProjectAdminOnly(projectRepository), projectController.DeleteProject)
}
