package cruds

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TTemplate struct {
	table string
}

var Template = &TTemplate{
	table: "template",
}

func (init *TTemplate) Get(id int) ([]*models.Template, error) {
	sqlString := `SELECT * FROM "template" WHERE id_template = $1;`

	data, err := database.BaseQuery[models.Template](sqlString, id)

	return data, err
}

func (init *TTemplate) GetAll() ([]*models.Template, error) {
	sqlString := `SELECT * FROM "template";`

	data, err := database.BaseQuery[models.Template](sqlString)

	return data, err
}

func (init *TTemplate) Create(rows [][]any) (int64, error) {
	columnNames := []string{"name", "description", "available"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TTemplate) Update(
	id int,
	name string,
	description string,
	available bool,
) ([]*models.Template, error) {
	sqlString := `
		UPDATE "template"
		SET (name, description, available)
		VALUES ($2, $3, $4)
		WHERE id_template = $1 
		RETURNING *;
	`
	args := []any{id, name, description, available}

	data, err := database.BaseQuery[models.Template](sqlString, args...)

	return data, err
}

func (init *TTemplate) Delete(id int) ([]*models.Template, error) {
	sqlString := `DELETE FROM "template" WHERE id_template = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Template](sqlString, id)

	return data, err
}

func (init *TTemplate) DeleteAll() ([]*models.Template, error) {
	sqlString := `DELETE FROM "template" RETURNING *;`

	data, err := database.BaseQuery[models.Template](sqlString)

	return data, err
}
