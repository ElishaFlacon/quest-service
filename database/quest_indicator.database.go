package database

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuestIndicator interface {
	Get() ([]models.QuestIndicator, error)
	GetAll() ([]models.QuestIndicator, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var QuestIndicator IQuestIndicator
