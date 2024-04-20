package cruds

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TQuest struct {
	table string
}

var Quest = &TQuest{
	table: "launch_quest",
}

func (init *TQuest) Get(id int) ([]*models.Quest, error) {
	sqlString := `SELECT * FROM "launch_quests" WHERE id_launch_quest = $1;`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return data, err
}

func (init *TQuest) GetAll() ([]*models.Quest, error) {
	sqlString := `SELECT * FROM "launch_quests";`

	data, err := database.BaseQuery[models.Quest](sqlString)

	return data, err
}

func (init *TQuest) Create(rows [][]any) (int64, error) {
	columnNames := []string{"available", "start_at", "end_at"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TQuest) Update(
	id int,
	idTeam int,
	available bool,
	startAt int,
	endAt int,
) ([]*models.Quest, error) {
	sqlString := `
		UPDATE "launch_quest" 
		SET (id_quest, id_team, available, start_at, end_at) 
		VALUES ($2, $3, $4, $5, $6) 
		WHERE id_quest=$1
		RETURNING *;
	`
	args := []any{id, idTeam, available, startAt, endAt}

	data, err := database.BaseQuery[models.Quest](sqlString, args...)

	return data, err
}

func (init *TQuest) Delete(id int) ([]*models.Quest, error) {
	sqlString := `DELETE * FROM "launch_quest" WHERE id_launch_quest = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return data, err
}

func (init *TQuest) DeleteAll() ([]*models.Quest, error) {
	sqlString := `DELETE * FROM "launch_quest" RETURNING *;`

	data, err := database.BaseQuery[models.Quest](sqlString)

	return data, err
}
