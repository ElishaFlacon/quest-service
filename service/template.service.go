package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TTemplate struct {
	table string
}

var Template = &TTemplate{
	table: "template",
}

func (*TTemplate) Get(id int) (*models.Template, error) {
	sqlString := `SELECT * FROM "template" WHERE id_template = $1;`

	data, errData := database.BaseQuery[models.Template](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	return data[0], nil
}

func (*TTemplate) GetWithIndicators(
	id int,
) (*models.TemplateWithIndicators, error) {
	template, errTemplate := Template.Get(id)
	if errTemplate != nil {
		return nil, errTemplate
	}

	indicators, errIndicators := Indicator.GetByTemplateId(id)
	if errIndicators != nil {
		return nil, errIndicators
	}

	newTemplate := &models.TemplateWithIndicators{
		IdTemplate:  template.IdTemplate,
		Name:        template.Name,
		Description: template.Description,
		Available:   template.Available,
		Indicators:  indicators,
	}

	return newTemplate, nil
}

func (*TTemplate) GetAll() ([]*models.Template, error) {
	sqlString := `SELECT * FROM "template";`

	data, errData := database.BaseQuery[models.Template](sqlString)
	if errData != nil {
		return nil, errData
	}

	return data, errData
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

	data, errData := database.BaseQuery[models.Template](
		sqlString,
		args...,
	)
	if errData != nil {
		return nil, 0, errData
	}

	var rows [][]any
	for index := range indicators {
		elememt := []any{data[0].IdTemplate, indicators[index]}
		rows = append(rows, elememt)
	}

	count, errTemplateIndicator := TemplateIndicator.CreateWithCopy(rows)
	if errTemplateIndicator != nil {
		return nil, 0, errTemplateIndicator
	}

	return data[0], count, nil
}

func (*TTemplate) Hide(id int) (*models.Template, error) {
	sqlString := `
		UPDATE "template" 
		SET available = false
		WHERE id_template = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Template](sqlString, id)

	return utils.CultivateFirstDataElement(data, err)
}
