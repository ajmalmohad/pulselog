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

	projectMemberRouter.POST("", projectMemberController.CreateProjectMember)
	projectMemberRouter.GET("/all", projectMemberController.GetAllProjectMembers) // Gets all project members by project id
	projectMemberRouter.GET("", middleware.SameProjectMemberOnly(projectMemberRepository), projectMemberController.GetProjectMember)
	projectMemberRouter.PUT("", middleware.SameProjectAdminOnly(projectMemberRepository), projectMemberController.UpdateProjectMember)
	projectMemberRouter.DELETE("", middleware.SameProjectAdminOnly(projectMemberRepository), projectMemberController.DeleteProjectMember)
}
