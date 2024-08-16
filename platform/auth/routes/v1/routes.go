package routes

import (
    "github.com/gin-gonic/gin"
    "your_project/controllers"
    "your_project/middlewares"
)

func SetupV1Routes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controllers.SignupHandler)
        auth.POST("/login", controllers.LoginHandler)
        auth.POST("/reauthenticate", middlewares.AuthMiddleware(), controllers.ReauthenticateHandler)
        auth.POST("/disable", middlewares.AuthMiddleware(), controllers.DisableUserHandler)
        auth.DELETE("/delete", middlewares.AuthMiddleware(), controllers.DeleteUserHandler)
    }
}
