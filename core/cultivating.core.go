package core

import "errors"

func CultivatingData[T comparable](data []T, err error) ([]T, error) {
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CultivatingOneData[T comparable](data []T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("empty data")
	}

	return &data[0], nil
}

// TODO: скорее всего стоит убрать из-за ненадобности
func CultivatingError(err error) error {
	if err != nil {
		return err
	}

	return nil
}
