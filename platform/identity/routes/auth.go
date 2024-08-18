package routes

import (
	"pulselog/identity/controllers"
	"pulselog/identity/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authRouter := router.Group("/auth")
	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)
	authController := controllers.NewAuthController(userRepository, refreshTokenRepository)

	authRouter.POST("/signup", authController.SignupHandler)
	authRouter.POST("/login", authController.LoginHandler)
	authRouter.POST("/reauthenticate", authController.ReauthenticateHandler)
}
