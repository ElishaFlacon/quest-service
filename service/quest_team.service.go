package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TQuestTeam struct {
	table string
}

var QuestTeam = &TQuestTeam{
	table: "quest_team",
}

func (*TQuestTeam) Get(id int) (*models.QuestTeam, error) {
	sqlString := `
		SELECT * 
		FROM "quest_team" 
		WHERE id_quest_team = $1;
	`

	data, err := database.BaseQuery[models.QuestTeam](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TQuestTeam) GetAll() ([]*models.QuestTeam, error) {
	sqlString := `SELECT * FROM "quest_team";`

	data, err := database.BaseQuery[models.QuestTeam](sqlString)

	return data, err
}

func (init *TQuestTeam) Create(rows [][]any) (int64, error) {
	columnNames := []string{"id_team", "id_quest"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (*TQuestTeam) Update(
	id int,
	id_team string,
	id_quest int,
) (*models.QuestTeam, error) {
	sqlString := `
		UPDATE "quest_team" 
		SET (id_team, id_quest)
		VALUES ($2, $3) 
		WHERE id_quest_team = $1
		RETURNING *;
	`
	args := []any{id, id_team, id_quest}

	data, err := database.BaseQuery[models.QuestTeam](
		sqlString,
		args...,
	)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TQuestTeam) Delete(id int) (*models.QuestTeam, error) {
	sqlString := `
		DELETE FROM "quest_team" 
		WHERE id_quest_team = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.QuestTeam](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TQuestTeam) DeleteAll() ([]*models.QuestTeam, error) {
	sqlString := `DELETE FROM "quest_team" RETURNING *;`

	data, err := database.BaseQuery[models.QuestTeam](sqlString)

	return data, err
}
