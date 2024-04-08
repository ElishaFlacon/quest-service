package service

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

func (*TIndicator) GetByQuestId(id int) ([]*models.Indicator, error) {
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

func (*TIndicator) GetByQuestAndUserId(idQuest int, idUser int) ([]*models.Indicator, error) {
	sqlString := `
		SELECT * FROM "indicator"
		INNER JOIN "template_indicator" ON "indicator".id_indicator = "template_indicator".id_template_indicator
		INNER JOIN "template" ON "template_indicator".id_template_indicator = "template".id_template
		INNER JOIN "quest" ON "template".id_template = "id_quest".id_quest
		INNER JOIN "result" "id_quest".id_quest = "result".id_quest
		WHERE id_quest = $1 AND id_from_user = $2
	`
	args := []any{idQuest, idUser}

	data, err := database.BaseQuery[models.Indicator](sqlString, args...)

	return data, err
}

func (*TIndicator) Create(name string, description string, role string, visible bool, idCategory int) ([]*models.Indicator, error) {
	sqlString := `
		INSERT INTO "indicator" 
		(name, description, role, visible, id_category) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *;
	`
	args := []any{name, description, role, true, idCategory}

	data, err := database.BaseQuery[models.Indicator](sqlString, args...)

	return data, err
}
