package routes

import (
	"pulselog/identity/controllers"
	"pulselog/identity/middleware"
	"pulselog/identity/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProjectMemberRoutes(router *gin.Engine, db *gorm.DB) {
	projectMemberRouter := router.Group("/project-members")
	userRepository := repositories.NewUserRepository(db)
	projectMemberRepository := repositories.NewProjectMemberRepository(db)
	projectMemberRouter.Use(middleware.AuthMiddleware(userRepository))
	projectMemberController := controllers.NewProjectMemberController(projectMemberRepository)

	projectMemberRouter.POST("", projectMemberController.CreateProjectMember)
	projectMemberRouter.GET("/all", projectMemberController.GetAllProjectMembers) // Gets all project members by project id
	// TODO: Add access control to this routes
	projectMemberRouter.GET("", projectMemberController.GetProjectMember)
	projectMemberRouter.PUT("", projectMemberController.UpdateProjectMember)
	projectMemberRouter.DELETE("", projectMemberController.DeleteProjectMember)
}
