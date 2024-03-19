package database

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IIndicator interface {
	Get() ([]models.Indicator, error)
	GetAll() ([]models.Indicator, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Indicator IIndicator
