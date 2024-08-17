package routes

import (
	"pulselog/auth/controllers"
	"pulselog/auth/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupV1Routes(router *gin.Engine, db *gorm.DB) {
	authRouter := router.Group("/v1/auth")
	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)
	authController := controllers.NewAuthController(
		userRepository,
		refreshTokenRepository,
	)
	{
		authRouter.POST("/signup", authController.SignupHandler)
		authRouter.POST("/login", authController.LoginHandler)
		authRouter.POST("/reauthenticate", authController.ReauthenticateHandler)
		authRouter.POST("/disable", authController.DisableUserHandler)
		authRouter.DELETE("/delete", authController.DeleteUserHandler)
	}
}
