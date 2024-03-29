package cruds

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuestIndicator struct {
	table string
}

var QuestIndicator = &TQuestIndicator{
	table: "quest_indicator",
}

func (init *TQuestIndicator) Get(id int) ([]*models.QuestIndicator, error) {
	sqlString := `SELECT * FROM "quest_indicator" WHERE idQuestIndicator = $1;`

	data, err := database.BaseQuery[models.QuestIndicator](sqlString, id)

	return data, err
}

func (init *TQuestIndicator) GetAll() ([]*models.QuestIndicator, error) {
	sqlString := `SELECT * FROM "quest_indicator";`

	data, err := database.BaseQuery[models.QuestIndicator](sqlString)

	return data, err
}

func (init *TQuestIndicator) Create(rows [][]any) (int64, error) {
	columnNames := []string{"id_quest", "id_indicator"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TQuestIndicator) Update(id_quest int, id_indicator int) ([]*models.QuestIndicator, error) {
	sqlString := `
		UPDATE "quest_indicator" 
		SET (id_quest, id_indicator) 
		VALUES ($1, $2) 
		WHERE id_quest_indicator = $1
		RETURNING *;
	`
	args := []any{id_quest, id_indicator}

	data, err := database.BaseQuery[models.QuestIndicator](sqlString, args...)

	return data, err
}

func (init *TQuestIndicator) Delete(id int) ([]*models.QuestIndicator, error) {
	sqlString := `DELETE * FROM "quest_indicator" WHERE id_quest_indicator = $1 RETURNING *;`

	data, err := database.BaseQuery[models.QuestIndicator](sqlString, id)

	return data, err
}

func (init *TQuestIndicator) DeleteAll() ([]*models.QuestIndicator, error) {
	sqlString := `DELETE * FROM "quest_indicator" RETURNING *;`

	data, err := database.BaseQuery[models.QuestIndicator](sqlString)

	return data, err
}
