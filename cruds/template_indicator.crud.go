package cruds

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TTemplateIndicator struct {
	table string
}

var TemplateIndicator = &TTemplateIndicator{
	table: "template_indicator",
}

func (init *TTemplateIndicator) Get(id int) ([]*models.TemplateIndicator, error) {
	sqlString := `SELECT * FROM "template_indicator" WHERE id_template_indicator = $1;`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString, id)

	return data, err
}

func (init *TTemplateIndicator) GetAll() ([]*models.TemplateIndicator, error) {
	sqlString := `SELECT * FROM "template_indicator";`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString)

	return data, err
}

func (init *TTemplateIndicator) Create(rows [][]any) (int64, error) {
	columnNames := []string{"id_quest", "id_indicator"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TTemplateIndicator) Update(id_quest int, id_indicator int) ([]*models.TemplateIndicator, error) {
	sqlString := `
		UPDATE "template_indicator" 
		SET (id_quest, id_indicator) 
		VALUES ($1, $2) 
		WHERE id_template_indicator = $1
		RETURNING *;
	`
	args := []any{id_quest, id_indicator}

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString, args...)

	return data, err
}

func (init *TTemplateIndicator) Delete(id int) ([]*models.TemplateIndicator, error) {
	sqlString := `DELETE * FROM "template_indicator" WHERE id_template_indicator = $1 RETURNING *;`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString, id)

	return data, err
}

func (init *TTemplateIndicator) DeleteAll() ([]*models.TemplateIndicator, error) {
	sqlString := `DELETE * FROM "template_indicator" RETURNING *;`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString)

	return data, err
}
