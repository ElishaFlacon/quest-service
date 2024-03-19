package database

import (
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type ILaunchQuest interface {
	Get() ([]models.LaunchQuest, error)
	GetAll() ([]models.LaunchQuest, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var LaunchQuest ILaunchQuest
