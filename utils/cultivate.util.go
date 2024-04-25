package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CultivateFirstDataElemet[T comparable](
	data []*T,
	err error,
) (*T, error) {
	if err != nil {
		return nil, err
	}

	return data[0], err
}

func CultivateCondition(
	context *gin.Context,
	condition bool,
	code int,
	message string,
) bool {
	if !condition {
		return false
	}

	context.JSON(code, gin.H{"error": message})
	context.Abort()
	return true
}

func CultivateBody[T comparable](
	context *gin.Context,
	object T,
) {
	err := context.BindJSON(object)

	if err == nil {
		return
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": err.Error()},
	)
	context.Abort()
}

func CultivateServiceError(
	context *gin.Context,
	err error,
) bool {
	if err == nil {
		return false
	}

	context.JSON(
		http.StatusInternalServerError,
		gin.H{"error": err.Error()},
	)
	context.Abort()
	return true
}

func CultivateServiceData(
	context *gin.Context,
	data any,
	err error,
) {
	isServiceError := CultivateServiceError(context, err)
	if isServiceError {
		return
	}

	context.JSON(http.StatusOK, data)
	context.Abort()
}

func CultivateParamsError(
	context *gin.Context,
	err error,
) {
	if err == nil {
		return
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": err.Error()},
	)
	context.Abort()
}

func CultivateStringParam(
	context *gin.Context,
	name string,
) string {
	param, errParam := GetStringParam(context.Params, name)
	CultivateParamsError(context, errParam)

	return param
}

func CultivateStringParams(
	context *gin.Context,
	names []string,
) []string {
	params, errParam := GetStringParams(context.Params, names)
	CultivateParamsError(context, errParam)

	return params
}

func CultivateNumberParam(
	context *gin.Context,
	name string,
) int {
	param, errParam := GetNumberParam(context.Params, name)
	CultivateParamsError(context, errParam)

	return param
}

func CultivateNumberParams(
	context *gin.Context,
	names []string,
) []int {
	params, errParam := GetNumberParams(context.Params, names)
	CultivateParamsError(context, errParam)

	return params
}
