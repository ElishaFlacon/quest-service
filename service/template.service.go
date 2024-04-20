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
	// TODO для Тимура: возвращаем все поля для шаблона
	return nil, nil
}

func (*TTemplate) GetWithIndicators(id int) (*models.TemplateWithIndicators, error) {
	// TODO для Тимура: возвращаем все поля для шаблона + массив вопросов (use indicators.GetByTemplateId)
	return nil, nil
}

func (*TTemplate) GetAll() ([]*models.Template, error) {
	// TODO для Тимура: возвращаем все шаблоны
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

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TTemplate) Delete(id int) (*models.Template, error) {
	sqlString := `
		DELETE FROM "template" 
		WHERE id_template = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Template](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}
