package service

// TODO add all methods

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

func (*TTemplate) Get(id int) (*models.Template, error) {
	// TODO

	return nil, nil
}

func (*TTemplate) GetAll() ([]*models.Template, error) {
	// TODO

	return nil, nil
}

func (*TTemplate) Create(
	name string,
	description string,
	indicators []int,
) (
	*models.Template,
	int64,
	error,
) {
	sqlString := `
		INSERT INTO "template" 
		(name, description) 
		VALUES ($1, $2) 
		RETURNING *;
	`
	args := []any{name, description}

	data, err := database.BaseQuery[models.Template](
		sqlString,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	tableName := "template_indicator"
	columnNames := []string{"id_template", "id_indicator"}
	var rows [][]any

	for index := range indicators {
		rows = append(
			rows,
			[]any{data[0].IdTemplate, indicators[index]},
		)
	}

	count, err := database.CopyFromQuery(
		tableName,
		columnNames,
		rows,
	)

	return data[0], count, err
}
