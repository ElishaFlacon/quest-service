package service

import "github.com/ElishaFlacon/quest-service/models"

type TCategory struct{}

var Category *TCategory

func (*TCategory) Get(id int) (*models.Category, error) {
	// TODO

	return nil, nil
}

func (*TCategory) GetAll() ([]*models.Category, error) {
	// TODO

	return nil, nil
}

func (*TCategory) Create(name string) (*models.Category, error) {
	// TODO

	return nil, nil
}

func (*TCategory) Hide(id int) (*models.Category, error) {
	// TODO

	return nil, nil
}

func (*TCategory) Delete(id int) (*models.Category, error) {
	// TODO

	return nil, nil
}
