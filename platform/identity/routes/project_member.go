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
	projectRepository := repositories.NewProjectRepository(db)
	projectMemberRepository := repositories.NewProjectMemberRepository(db)
	projectMemberRouter.Use(middleware.AuthMiddleware(userRepository))
	projectMemberController := controllers.NewProjectMemberController(projectRepository, projectMemberRepository)

	// TODO: Restrict people from creating duplicate project members
	projectMemberRouter.POST("", projectMemberController.CreateProjectMember)     // Only admins are able to create project members
	projectMemberRouter.GET("/all", projectMemberController.GetAllProjectMembers) // Gets all project members by project id
	projectMemberRouter.GET("", middleware.SameProjectMemberOnly(projectMemberRepository), projectMemberController.GetProjectMember)
	projectMemberRouter.PUT("", middleware.SameProjectAdminOnly(projectMemberRepository), projectMemberController.UpdateProjectMember)
	projectMemberRouter.DELETE("", middleware.SameProjectAdminOnly(projectMemberRepository), projectMemberController.DeleteProjectMember)
}
