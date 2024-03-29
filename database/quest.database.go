package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IQuest interface {
	GetQuest(id int) ([]models.Quest, error)
	GetAllQuests() ([]models.Quest, error)
	CreateQuest(quest models.Quest) ([]*models.Quest, error)
	UpdateQuest(rows [][]any) (int64, error)
	DeleteQuest(id int) ([]*models.Quest, error)
	DeleteAllQuests() ([]*models.Quest, error)
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

func (repo *TDatabase) CreateQuest(quest models.Quest) ([]*models.Quest, error) {
	sqlString := `INSERT INTO quest (available, name, description, link) VALUES ($1, $2, $3, $4) RETURNING *;`

	data, err := core.QueryWithReturningData[models.Quest](
		sqlString,
		repo.db.Query,
		quest.Available,
		quest.Name,
		quest.Description,
		quest.Link,
	)
	return data, err
}

func (repo *TDatabase) UpdateQuest(rows [][]any) (int64, error) {
	tableName := "quest"
	columnNames := []string{"available", "name", "description", "link"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) DeleteQuest(id int) ([]*models.Quest, error) {
	sqlString := `DELETE FROM quest WHERE id_quest = $1 RETURNING *;`

	data, err := core.QueryWithReturningData[models.Quest](sqlString, repo.db.Query, id)
	return data, err
}

func (repo *TDatabase) DeleteAllQuests() ([]*models.Quest, error) {
	sqlString := `DELETE FROM quest RETURNING *;`

	data, err := core.QueryWithReturningData[models.Quest](sqlString, repo.db.Query)
	return data, err
}
