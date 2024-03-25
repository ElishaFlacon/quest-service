package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuestIndicator interface {
	GetQuestIndicator(id int) ([]models.QuestIndicator, error)
	GetAllQuestIndicators() ([]models.QuestIndicator, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var QuestIndicator IQuestIndicator

func (repo *TDatabase) GetQuestIndicator(id int) ([]*models.QuestIndicator, error) {
	sqlString := `SELECT * FROM quest_indicator WHERE idQuestIndicator = $1;`

	data, err := core.QueryWithReturningData[models.QuestIndicator](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllQuestIndicators() ([]*models.QuestIndicator, error) {
	sqlString := `SELECT * FROM quest_indicator;`

	data, err := core.QueryWithReturningData[models.QuestIndicator](sqlString, repo.db.Query)

	return data, err
}
