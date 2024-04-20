package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errParam = errors.New("request params error")

func GetStringParam(
	params gin.Params,
	name string,
) (string, error) {
	strParam, paramGood := params.Get(name)

	if !paramGood {
		return "", errParam
	}

	return strParam, nil
}

func GetStringParams(
	params gin.Params,
	names []string,
) ([]string, error) {
	result := []string{}

	for index := range names {
		name := names[index]

		strParam, errParam := GetStringParam(params, name)

		if errParam != nil {
			return []string{}, errParam
		}

		result = append(result, strParam)
	}

	return result, nil
}

func GetNumberParam(
	params gin.Params,
	name string,
) (int, error) {
	strParam, paramGood := params.Get(name)

	if !paramGood {
		return 0, errParam
	}

	numberParam, errAtoi := strconv.Atoi(strParam)

	if errAtoi != nil {
		return 0, errParam
	}

	return numberParam, nil
}

func GetNumberParams(
	params gin.Params,
	names []string,
) ([]int, error) {
	result := []int{}

	for index := range names {
		name := names[index]

		numberParam, errParam := GetNumberParam(params, name)

		if errParam != nil {
			return []int{}, errParam
		}

		result = append(result, numberParam)
	}

	return result, nil
}
