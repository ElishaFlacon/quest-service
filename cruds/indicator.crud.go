package cruds

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TIndicator struct {
	table string
}

var Indicator = &TIndicator{
	table: "indicator",
}

func (init *TIndicator) Get(id int) ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM "indicator" WHERE id_indicator = $1;`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return data, err
}

func (init *TIndicator) GetAll() ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM "indicator";`

	data, err := database.BaseQuery[models.Indicator](sqlString)

	return data, err
}

func (init *TIndicator) Create(rows [][]any) (int64, error) {
	columnNames := []string{"name", "description", "role", "visible", "id_category"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TIndicator) Update(
	id int,
	name string,
	description string,
	role string,
	visible bool,
) ([]*models.Indicator, error) {
	sqlString := `
		UPDATE "indicator" 
		SET (name, description, role, visible) 
		VALUES ($2, $3, $4, $5)
		WHERE id_indicator=$1 
		RETURNING *;
	`
	args := []any{id, name, description, role, visible}

	data, err := database.BaseQuery[models.Indicator](sqlString, args...)

	return data, err
}

func (init *TIndicator) Delete(id int) ([]*models.Indicator, error) {
	sqlString := `DELETE * FROM "indicator" WHERE id_indicator = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return data, err
}

func (init *TIndicator) DeleteAll() ([]*models.Indicator, error) {
	sqlString := `DELETE * FROM "indicator" RETURNING *;`

	data, err := database.BaseQuery[models.Indicator](sqlString)

	return data, err
}
