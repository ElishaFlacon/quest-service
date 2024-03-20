package database

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuest interface {
	Get() ([]models.Quest, error)
	GetAll() ([]models.Quest, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Quest IQuest
