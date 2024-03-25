package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type ILaunchQuest interface {
	GetLaunchQuest(id int) ([]models.LaunchQuest, error)
	GetAllLaunchQuests() ([]models.LaunchQuest, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var LaunchQuest ILaunchQuest

func (repo *TDatabase) GetLaunchQuest(id int) ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM launch_quests WHERE inLaunchIndicator = $1;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllLaunchQuests() ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM launch_quests;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query)

	return data, err
}
