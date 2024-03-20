package utils

func NormalizeData[T comparable](data []T, err error) ([]T, error) {
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NormalizeError(err error) error {
	if err != nil {
		return err
	}

	return nil
}
