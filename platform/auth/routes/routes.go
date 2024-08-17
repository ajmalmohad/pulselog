package routes

import (
	"pulselog/auth/controllers"
	"pulselog/auth/middleware"
	"pulselog/auth/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	authRouter := router.Group("/auth")
	userRouter := router.Group("/users")
	userRouter.Use(middleware.AuthMiddleware())

	userRepository, refreshTokenRepository := initializeRepositories(db)
	authController, userController := initializeControllers(userRepository, refreshTokenRepository)

	setupAuthRoutes(authRouter, authController)
	setupUserRoutes(userRouter, userController)
}

func initializeRepositories(db *gorm.DB) (*repositories.UserRepository, *repositories.RefreshTokenRepository) {
	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)
	return userRepository, refreshTokenRepository
}

func initializeControllers(userRepo *repositories.UserRepository, refreshTokenRepo *repositories.RefreshTokenRepository) (*controllers.AuthController, *controllers.UserController) {
	authController := controllers.NewAuthController(userRepo, refreshTokenRepo)
	userController := controllers.NewUserController(userRepo)
	return authController, userController
}

func setupAuthRoutes(router *gin.RouterGroup, controller *controllers.AuthController) {
	router.POST("/signup", controller.SignupHandler)
	router.POST("/login", controller.LoginHandler)
	router.POST("/reauthenticate", controller.ReauthenticateHandler)
}

func setupUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	router.DELETE("/delete", controller.DeleteUserHandler)
}
