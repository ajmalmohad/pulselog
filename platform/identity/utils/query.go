package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetValueFromQuery(ctx *gin.Context, key string, out interface{}) error {
	valueStr := ctx.Query(key)
	if valueStr == "" {
		return fmt.Errorf("%s is required", key)
	}

	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr || outValue.IsNil() {
		return fmt.Errorf("out parameter must be a non-nil pointer")
	}

	outElem := outValue.Elem()
	switch outElem.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		parsedValue, err := strconv.ParseUint(valueStr, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid %s", key)
		}
		outElem.SetUint(parsedValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parsedValue, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid %s", key)
		}
		outElem.SetInt(parsedValue)
	case reflect.String:
		outElem.SetString(valueStr)
	default:
		return fmt.Errorf("unsupported type for %s", key)
	}

	return nil
}

func GetProjectIDFromQuery(ctx *gin.Context) (uint, error) {
	var projectID uint
	err := GetValueFromQuery(ctx, "project_id", &projectID)
	if err != nil {
		return 0, err
	}
	return projectID, nil
}

func GetProjectMemberIDFromQuery(ctx *gin.Context) (uint, error) {
	var projectMemberID uint
	err := GetValueFromQuery(ctx, "project_member_id", &projectMemberID)
	if err != nil {
		return 0, err
	}
	return projectMemberID, nil
}

func GetAPIKeyIDFromQuery(ctx *gin.Context) (uint, error) {
	var apiKeyID uint
	err := GetValueFromQuery(ctx, "api_key_id", &apiKeyID)
	if err != nil {
		return 0, err
	}
	return apiKeyID, nil
}
