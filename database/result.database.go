package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IResult interface {
	GetResult(id int) ([]models.Result, error)
	GetAllResults() ([]models.Result, error)
	CreateResult(result models.Result) ([]*models.Result, error)
	UpdateResult(rows [][]any) (int64, error)
	DeleteResult(id int) ([]*models.Result, error)
	DeleteAllResults() ([]*models.Result, error)
}

var Result IResult

func (repo *TDatabase) GetResult(id int) ([]*models.Result, error) {
	sqlString := `SELECT * FROM result WHERE idResult = $1;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllResults() ([]*models.Result, error) {
	sqlString := `SELECT * FROM result;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query)

	return data, err
}

func (repo *TDatabase) CreateResult(result models.Result) ([]*models.Result, error) {
	sqlString := `INSERT INTO result (id_indicator, id_launch_quest, id_from_user, id_to_user, value) VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	data, err := core.QueryWithReturningData[models.Result](
		sqlString,
		repo.db.Query,
		result.IdIndicator,
		result.IdLaunchQuest,
		result.IdFromUser,
		result.IdToUser,
		result.Value,
	)
	return data, err
}

func (repo *TDatabase) UpdateResult(rows [][]any) (int64, error) {
	tableName := "result"
	columnNames := []string{"id_indicator", "id_launch_quest", "id_from_user", "id_to_user", "value"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) DeleteResult(id int) ([]*models.Result, error) {
	sqlString := `DELETE FROM result WHERE id_result = $1 RETURNING *;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query, id)
	return data, err
}

func (repo *TDatabase) DeleteAllResults() ([]*models.Result, error) {
	sqlString := `DELETE FROM result RETURNING *;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query)
	return data, err
}
