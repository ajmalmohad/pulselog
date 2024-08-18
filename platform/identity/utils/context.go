package utils

import (
	"fmt"
	"net/http"
	"pulselog/identity/types"

	"github.com/gin-gonic/gin"
)

func InjectClaimsToContext(ctx *gin.Context, tokenString string) {
	userID, email, err := ExtractUserIDAndEmailFromClaims(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:  "Failed to extract user ID and email from claims",
			Detail: err.Error(),
		})
		return
	}

	ctx.Set("user_id", userID)
	ctx.Set("email", email)
}

func ExtractClaimsFromContext(ctx *gin.Context) (uint, string, error) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, "", fmt.Errorf("user ID not found in context")
	}
	email, exists := ctx.Get("email")
	if !exists {
		return 0, "", fmt.Errorf("email not found in context")
	}

	uid, ok := userID.(uint)
	if !ok {
		return 0, "", fmt.Errorf("user ID has an invalid type")
	}
	eml, ok := email.(string)
	if !ok {
		return 0, "", fmt.Errorf("email has an invalid type")
	}

	return uid, eml, nil
}
