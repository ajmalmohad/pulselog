package middleware

import (
	"net/http"
	"pulselog/identity/repositories"
	"pulselog/identity/types"
	"pulselog/identity/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(
	userRepository *repositories.UserRepository,
) gin.HandlerFunc {
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

		userID, _, err := utils.ExtractUserIDAndEmailFromClaims(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error:  "Failed to extract user ID and email from claims",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		// Check if user exists in the database
		_, err = userRepository.FindByID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error:  "Failed to find user by ID",
				Detail: err.Error(),
			})
			ctx.Abort()
			return
		}

		utils.InjectClaimsToContext(ctx, tokenString)

		ctx.Next()
	}
}
