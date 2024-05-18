package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// тут много всяких полезных обрабатывалок для controllers
// но лучше не залазьте сюда, душный файл

func CultivateFirstDataElemet[T any](
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

func CultivateBody[T any](
	context *gin.Context,
	object T,
) error {
	err := context.BindJSON(object)

	if err == nil {
		return nil
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": err.Error()},
	)
	context.Abort()

	return err
}

func CultivateServiceError(
	context *gin.Context,
	err error,
) error {
	if err == nil {
		return nil
	}

	context.JSON(
		http.StatusInternalServerError,
		gin.H{"error": err.Error()},
	)
	context.Abort()

	return err
}

func CultivateServiceData(
	context *gin.Context,
	data any,
	err error,
) error {
	errCultivate := CultivateServiceError(context, err)
	if errCultivate != nil {
		return errCultivate
	}

	context.JSON(http.StatusOK, data)
	context.Abort()

	return nil
}

func CultivateParamsError(
	context *gin.Context,
	err error,
) error {
	if err == nil {
		return nil
	}

	context.JSON(
		http.StatusBadRequest,
		gin.H{"error": err.Error()},
	)
	context.Abort()

	return err
}

func CultivateStringParam(
	context *gin.Context,
	name string,
) (string, error) {
	param, errParam := GetStringParam(context.Params, name)
	CultivateParamsError(context, errParam)

	return param, errParam
}

func CultivateStringParams(
	context *gin.Context,
	names []string,
) ([]string, error) {
	params, errParam := GetStringParams(context.Params, names)
	CultivateParamsError(context, errParam)

	return params, errParam
}

func CultivateNumberParam(
	context *gin.Context,
	name string,
) (int, error) {
	param, errParam := GetNumberParam(context.Params, name)
	CultivateParamsError(context, errParam)

	return param, errParam
}

func CultivateNumberParams(
	context *gin.Context,
	names []string,
) ([]int, error) {
	params, errParam := GetNumberParams(context.Params, names)
	CultivateParamsError(context, errParam)

	return params, errParam
}
