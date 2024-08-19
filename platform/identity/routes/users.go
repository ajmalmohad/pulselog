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
	userRouter.Use(middleware.AuthMiddleware())
	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	userRouter.DELETE("", userController.DeleteUserHandler)
}
