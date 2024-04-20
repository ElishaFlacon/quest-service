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
) {
	if !condition {
		return
	}

	context.JSON(code, gin.H{"error": message})
	context.Abort()
}

func CultivateBody(
	context *gin.Context,
	object any,
) {
	err := context.BindJSON(&object)

	if err == nil {
		return
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": errParam.Error()},
	)
	context.Abort()
}

func CultivateId(data []struct{ Id int }) []int {
	result := []int{}

	for index := range data {
		result = append(result, data[index].Id)
	}

	return result
}

func CultivateServiceError(
	context *gin.Context,
	err error,
) {
	if err == nil {
		return
	}

	context.JSON(
		http.StatusInternalServerError,
		gin.H{"error": err.Error()},
	)
	context.Abort()
}

func CultivateServiceData(
	context *gin.Context,
	data any,
	err error,
) {
	CultivateServiceError(context, err)
	context.JSON(http.StatusOK, data)
	context.Abort()
}

func CultivateParamsError(
	context *gin.Context,
	err error,
) {
	if errParam == nil {
		return
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": errParam.Error()},
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
