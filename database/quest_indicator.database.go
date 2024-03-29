package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuestIndicator interface {
	GetQuestIndicator(id int) ([]models.QuestIndicator, error)
	GetAllQuestIndicators() ([]models.QuestIndicator, error)
	CreateQuestIndicator(questIndicator models.QuestIndicator) ([]*models.QuestIndicator, error)
	UpdateQuestIndicator(rows [][]any) (int64, error)
	DeleteQuestIndicator(id int) ([]*models.QuestIndicator, error)
	DeleteAllQuestIndicators() ([]*models.QuestIndicator, error)
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

func (repo *TDatabase) CreateQuestIndicator(questIndicator models.QuestIndicator) ([]*models.QuestIndicator, error) {
	sqlString := `INSERT INTO quest_indicator (id_quest, id_indicator) VALUES ($1, $2) RETURNING *;`

	data, err := core.QueryWithReturningData[models.QuestIndicator](
		sqlString,
		repo.db.Query,
		questIndicator.IdQuest,
		questIndicator.IdIndicator,
	)
	return data, err
}

func (repo *TDatabase) UpdateQuestIndicator(rows [][]any) (int64, error) {
	tableName := "quest_indicator"
	columnNames := []string{"id_quest", "id_indicator"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) DeleteQuestIndicator(id int) ([]*models.QuestIndicator, error) {
	sqlString := `DELETE FROM quest_indicator WHERE id_quest_indicator = $1 RETURNING *;`

	data, err := core.QueryWithReturningData[models.QuestIndicator](sqlString, repo.db.Query, id)
	return data, err
}

func (repo *TDatabase) DeleteAllQuestIndicators() ([]*models.QuestIndicator, error) {
	sqlString := `DELETE FROM quest_indicator RETURNING *;`

	data, err := core.QueryWithReturningData[models.QuestIndicator](sqlString, repo.db.Query)
	return data, err
}
