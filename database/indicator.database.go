package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

const table = "indicator"

type IIndicator interface {
	GetIndicator(id int) ([]*models.Indicator, error)
	GetAllIndicators() ([]*models.Indicator, error)
	CreateIndicator(indicator models.Indicator) ([]models.Indicator, error)
	UpdateIndicator(row [][]any) (int64, error)
	DeleteIndicator(id int) ([]*models.Indicator, error)
	DeleteAllIndicators() ([]*models.Indicator, error)
}

var Indicator IIndicator

func (repo *TDatabase) GetIndicator(id int) ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM indicator WHERE id_indicator = $1;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllIndicators() ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM indicator;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query)

	return data, err
}

func (repo *TDatabase) CreateIndicator(indicator models.Indicator) ([]*models.Indicator, error) {
	sqlString := `INSERT INTO indicator (value, description, type, role, visible) VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	data, err := core.QueryWithReturningData[models.Indicator](
		sqlString,
		repo.db.Query,
		indicator.Value,
		indicator.Description,
		indicator.Type,
		indicator.Role,
		indicator.Visible,
	)
	return data, err
}

func (repo *TDatabase) UpdateIndicator(rows [][]any) (int64, error) {
	tableName := "indicator"
	columnNames := []string{"value", "description", "type", "role", "visible"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) DeleteIndicator(id int) ([]*models.Indicator, error) {
	sqlString := `DELETE FROM indicator WHERE id_indicator = $1 RETURNING *;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query, id)
	return data, err
}

func (repo *TDatabase) DeleteAllIndicators(id int) ([]*models.Indicator, error) {
	sqlString := `DELETE FROM indicator RETURNING *;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query, id)
	return data, err
}
