package routes

import (
	"pulselog/identity/controllers"
	"pulselog/identity/middleware"
	"pulselog/identity/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRouter := router.Group("/users")
	userRepository := repositories.NewUserRepository(db)
	userRouter.Use(middleware.AuthMiddleware(userRepository))
	userController := controllers.NewUserController(userRepository)

	userRouter.DELETE("", userController.DeleteUserHandler)
	userRouter.DELETE("/logout", userController.LogoutUserHandler)
	userRouter.DELETE("/logout/all", userController.LogoutAllUserHandler)
}
