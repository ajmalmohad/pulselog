package middleware

import (
	"net/http"
	"pulselog/identity/types"
	"pulselog/identity/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{
				Error: "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		var tokenString string
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenString = authHeader
		}

		err := utils.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{
				Error:  "Invalid token",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		userID, email, err := utils.ExtractUserIDAndEmailFromClaims(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error:  "Failed to extract user ID and email from claims",
				Detail: err.Error(),
			})
			return
		}

		ctx.Set("user_id", userID)
		ctx.Set("email", email)

		ctx.Next()
	}
}
