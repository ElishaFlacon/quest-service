package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TTemplate struct {
	table string
}

var Template = &TTemplate{
	table: "template",
}

func (*TIndicator) GetWithIndicators(id int) ([]*models.Indicator, error) {
	sqlString := `
		SELECT "indicator".id_indicator, "indicator".name, "indicator".description, "indicator".role, "indicator".visible
		FROM "indicator"
		INNER JOIN "template_indicator" ON "indicator".id_indicator = "template_indicator".id_indicator
		INNER JOIN "template" ON "template_indicator".id_template = "template".id_template
		INNER JOIN "quest" ON "template".id_template = "quest".id_template
		WHERE "quest".id_quest = $1
	`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return data, err
}

func (*TTemplate) Create(name string, description string, indicators []int) ([]*models.Template, int64, error) {
	sqlString := `INSERT INTO "template" (name, description) VALUES ($1, $2) RETURNING *;`
	args := []any{name, description}

	data, err := database.BaseQuery[models.Template](sqlString, args...)

	if err != nil {
		return nil, 0, err
	}

	tableName := "template_indicator"
	columnNames := []string{"id_template", "id_indicator"}
	var rows [][]any

	for index := 0; len(indicators) > index; index++ {
		rows = append(rows, []any{data[0].IdTemplate, indicators[index]})
	}

	count, err := database.CopyFromQuery(tableName, columnNames, rows)

	return data, count, err
}
