package middleware

import (
	"net/http"
	"pulselog/auth/types"
	"pulselog/auth/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Authorization header is required"})
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
			ctx.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Invalid token"})
			ctx.Abort()
			return
		}

		claims, err := utils.ExtractClaims(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract claims"})
			ctx.Abort()
			return
		}

		userID, ok := claims["user_id"].(uint)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract user ID from claims"})
			ctx.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Failed to extract email from claims"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)
		ctx.Set("email", email)

		ctx.Next()
	}
}
