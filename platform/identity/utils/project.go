package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProjectIDFromQuery(ctx *gin.Context) (uint, error) {
	projectIDStr := ctx.Query("project_id")
	if projectIDStr == "" {
		return 0, fmt.Errorf("project id is required")
	}

	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid project id")
	}

	return uint(projectID), nil
}
