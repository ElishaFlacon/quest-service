package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuest interface {
	GetQuest(id int) ([]models.Quest, error)
	GetAllQuests() ([]models.Quest, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Quest IQuest

func (repo *TDatabase) GetQuest(id int) ([]*models.Quest, error) {
	sqlString := `SELECT * FROM quest WHERE idQuest = $1;`

	data, err := core.QueryWithReturningData[models.Quest](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllQuests() ([]*models.Quest, error) {
	sqlString := `SELECT * FROM quest;`

	data, err := core.QueryWithReturningData[models.Quest](sqlString, repo.db.Query)

	return data, err
}
