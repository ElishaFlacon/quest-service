package cruds

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
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
	columnNames := []string{"value", "description", "type", "role", "visible"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TIndicator) Update(
	id int,
	value string,
	description string,
	itype string,
	role string,
	visible bool,
) ([]*models.Indicator, error) {
	sqlString := `
		UPDATE "indicator" 
		SET (value, description, type, role, visible) 
		VALUES ($2, $3, $4, $5, $6)
		WHERE id_indicator=$1 
		RETURNING *;
	`
	args := []any{id, value, description, itype, role, visible}

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
