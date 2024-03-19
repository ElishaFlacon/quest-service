package database

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IResult interface {
	Get() ([]models.Result, error)
	GetAll() ([]models.Result, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Result IResult
