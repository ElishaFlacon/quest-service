package cruds

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TLaunchQuest struct {
	table string
}

var LaunchQuest = &TLaunchQuest{
	table: "launch_quest",
}

func (init *TLaunchQuest) Get(id int) ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM "launch_quests" WHERE id_launch_quest = $1;`

	data, err := database.BaseQuery[models.LaunchQuest](sqlString, id)

	return data, err
}

func (init *TLaunchQuest) GetAll() ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM "launch_quests";`

	data, err := database.BaseQuery[models.LaunchQuest](sqlString)

	return data, err
}

func (init *TLaunchQuest) Create(
	id int,
	idTeam int,
	available bool,
	startAt int,
	endAt int,
) ([]*models.LaunchQuest, error) {
	sqlString := `
		UPDATE "launch_quest" 
		SET (id_quest, id_team, available, start_at, end_at) 
		VALUES ($2, $3, $4, $5, $6) 
		WHERE id_quest=$1
		RETURNING *;
	`
	args := []any{id, idTeam, available, startAt, endAt}

	data, err := database.BaseQuery[models.LaunchQuest](sqlString, args...)

	return data, err
}

func (init *TLaunchQuest) Update(rows [][]any) (int64, error) {
	columnNames := []string{"available", "start_at", "end_at"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TLaunchQuest) Delete(id int) ([]*models.LaunchQuest, error) {
	sqlString := `DELETE * FROM "launch_quest" WHERE id_launch_quest = $1 RETURNING *;`

	data, err := database.BaseQuery[models.LaunchQuest](sqlString, id)

	return data, err
}

func (init *TLaunchQuest) DeleteAll() ([]*models.LaunchQuest, error) {
	sqlString := `DELETE * FROM "launch_quest" RETURNING *;`

	data, err := database.BaseQuery[models.LaunchQuest](sqlString)

	return data, err
}
