package service

import (
	"github.com/ElishaFlacon/questionnaire-service/cruds"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TIndicator struct {
	table string
}

var Indicator = &TIndicator{
	table: "indicator",
}

func (*TIndicator) GetAll() ([]*models.IndicatorResponse, error) {
	sqlString := `SELECT id_indicator, name, description FROM "indicator";`

	data, err := database.BaseQuery[models.IndicatorResponse](sqlString)

	return data, err
}

func (*TIndicator) GetByQuestId(id int) ([]*models.IndicatorResponse, error) {
	sqlString := `
		SELECT "indicator".id_indicator, "indicator".name, "indicator".description FROM "indicator"
		INNER JOIN "template_indicator" ON "indicator".id_indicator = "template_indicator".id_indicator
		INNER JOIN "template" ON "template_indicator".id_template = "template".id_template
		INNER JOIN "quest" ON "template".id_template = "quest".id_template
		WHERE "quest".id_quest = $1
	`

	data, err := database.BaseQuery[models.IndicatorResponse](sqlString, id)

	return data, err
}

func (*TIndicator) GetByQuestAndUserId(idQuest int, idUser int) ([]*models.IndicatorResponse, error) {
	sqlString := `
		SELECT * FROM "indicator"
		INNER JOIN "template_indicator" ON "indicator".id_indicator = "template_indicator".id_template_indicator
		INNER JOIN "template" ON "template_indicator".id_template_indicator = "template".id_template
		INNER JOIN "quest" ON "template".id_template = "id_quest".id_quest
		INNER JOIN "result" "id_quest".id_quest = "result".id_quest
		WHERE id_quest = $1 AND id_from_user = $2
	`
	args := []any{idQuest, idUser}

	data, err := database.BaseQuery[models.IndicatorResponse](sqlString, args...)

	return data, err
}

func (*TIndicator) Create(name string, description string, IdCategory int) (*models.IndicatorResponse, error) {
	args := [][]any{{IdCategory, name, description}}

	count, err := cruds.Indicator.Create(args)

	if count == 0 {
		return nil, err
	}

	data := &models.IndicatorResponse{
		Name:        name,
		Description: description,
	}

	return data, err
}
