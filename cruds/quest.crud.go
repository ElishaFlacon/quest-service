package cruds

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuest struct {
	table string
}

var Quest = &TQuest{
	table: "quest",
}

func (init *TQuest) Get(id int) ([]*models.Quest, error) {
	sqlString := `SELECT * FROM "quest" WHERE id_quest = $1;`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return data, err
}

func (init *TQuest) GetAll() ([]*models.Quest, error) {
	sqlString := `SELECT * FROM "quest";`

	data, err := database.BaseQuery[models.Quest](sqlString)

	return data, err
}

func (init *TQuest) Create(rows [][]any) (int64, error) {
	columnNames := []string{"name", "description", "available", "start_at", "end_at"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TQuest) Update(
	id int,
	name string,
	description string,
	available bool,
	startAt int,
	endAt int,
) ([]*models.Quest, error) {
	sqlString := `
		UPDATE "quest"
		SET (name, description, available, start_at, end_at)
		VALUES ($2, $3, $4, $5, $6)
		WHERE id_quest = $1 
		RETURNING *;
	`
	args := []any{id, name, description, available, startAt, endAt}

	data, err := database.BaseQuery[models.Quest](sqlString, args...)

	return data, err
}

func (init *TQuest) Delete(id int) ([]*models.Quest, error) {
	sqlString := `DELETE FROM "quest" WHERE id_quest = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return data, err
}

func (init *TQuest) DeleteAll() ([]*models.Quest, error) {
	sqlString := `DELETE FROM "quest" RETURNING *;`

	data, err := database.BaseQuery[models.Quest](sqlString)

	return data, err
}
