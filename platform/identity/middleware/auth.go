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

		utils.InjectClaimsToContext(ctx, tokenString)

		ctx.Next()
	}
}
