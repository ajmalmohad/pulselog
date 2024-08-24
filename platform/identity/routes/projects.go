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
	userRepository := repositories.NewUserRepository(db)
	projectRepository := repositories.NewProjectRepository(db)
	projectMemberRepository := repositories.NewProjectMemberRepository(db)
	projectRouter.Use(middleware.AuthMiddleware(userRepository))
	projectController := controllers.NewProjectController(projectRepository, projectMemberRepository)

	projectRouter.POST("", projectController.CreateProject)
	projectRouter.GET("/all", projectController.GetAllProjects) // Gets all projects that the user is a member of
	projectRouter.GET("", middleware.ProjectMemberOnly(projectRepository), projectController.GetProject)
	projectRouter.PUT("", middleware.ProjectAdminOnly(projectRepository), projectController.UpdateProject)
	projectRouter.DELETE("", middleware.ProjectAdminOnly(projectRepository), projectController.DeleteProject)
}
