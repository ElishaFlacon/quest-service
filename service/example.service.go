package service

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

func (init *TInit) GetExample() ([]models.Example, error) {
	data, err := init.db.GetExample()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (init *TInit) SetExample(value string) error {
	err := init.db.SetExample(value)

	if err != nil {
		return err
	}

	return nil
}
