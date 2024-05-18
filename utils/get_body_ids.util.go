package utils

import "github.com/ElishaFlacon/quest-service/models"

func GetBodyIds[T any](data []models.Id[T]) []T {
	result := []T{}
	for index := range data {
		result = append(result, data[index].Id)
	}

	return result
}
