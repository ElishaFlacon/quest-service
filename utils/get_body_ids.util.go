package utils

func GetBodyIds[T comparable](data []struct{ Id T }) []T {
	result := []T{}

	for index := range data {
		result = append(result, data[index].Id)
	}

	return result
}
