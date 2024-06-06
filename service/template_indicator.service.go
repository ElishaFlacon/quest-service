package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TTemplateIndicator struct {
	table string
}

var TemplateIndicator = &TTemplateIndicator{
	table: "template_indicator",
}

func (init *TTemplateIndicator) Get(id int) (*models.TemplateIndicator, error) {
	sqlString := `
		SELECT * FROM "template_indicator" 
		WHERE id_template_indicator = $1;
	`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString, id)

	return utils.CultivateFirstDataElement(data, err)
}

func (init *TTemplateIndicator) GetAll() ([]*models.TemplateIndicator, error) {
	sqlString := `SELECT * FROM "template_indicator";`

	data, err := database.BaseQuery[models.TemplateIndicator](sqlString)

	return data, err
}

func (init *TTemplateIndicator) CreateWithCopy(rows [][]any) (int64, error) {
	columnNames := []string{"id_template", "id_indicator"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}
