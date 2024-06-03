package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/jackc/pgx/v5"
)

type TQuestTeam struct {
	table string
}

var QuestTeam = &TQuestTeam{
	table: "quest_team",
}

func (*TQuestTeam) Get(id int) (*models.QuestTeam, error) {
	sqlString := `
		SELECT * FROM "quest_team" 
		WHERE id_quest_team = $1;
	`

	data, err := database.BaseQuery[models.QuestTeam](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TQuestTeam) GetByQuestId(id int) ([]*models.QuestTeam, error) {
	sqlString := `
		SELECT 
			"quest_team".id_quest_team,
			"quest_team".id_quest,
			"quest_team".id_team,
			"quest_team".name
		FROM "quest_team"
		INNER JOIN "quest" ON 
			"quest_team".id_quest = "quest".id_quest
		WHERE "quest".id_quest = $1;
	`

	data, err := database.BaseQuery[models.QuestTeam](sqlString, id)

	return data, err
}

func (*TQuestTeam) GetAll() ([]*models.QuestTeam, error) {
	sqlString := `SELECT * FROM "quest_team";`

	data, err := database.BaseQuery[models.QuestTeam](sqlString)

	return data, err
}

func (init *TQuestTeam) CreateWithCopy(rows [][]any) (int64, error) {
	columnNames := []string{"id_quest", "id_team"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (*TQuestTeam) CreateWithBatch(
	idQuest int,
	teams []*models.Team,
) ([]*models.QuestTeam, error) {
	sqlString := `
		INSERT INTO "quest_team"
		(id_quest, id_team, name)
		VALUES ($1, $2, $3)
		RETURNING *;
	`

	batch := &pgx.Batch{}

	for _, team := range teams {
		batch.Queue(sqlString, idQuest, team.IdTeam, team.Name)
	}

	data, err := database.SendBatch[models.QuestTeam](batch)

	return data, err
}

func (*TQuestTeam) Update(
	id int,
	id_quest int,
	id_team string,
	name string,
) (*models.QuestTeam, error) {
	sqlString := `
		UPDATE "quest_team" 
		SET (id_quest, id_team)
		VALUES ($2, $3, $4) 
		WHERE id_quest_team = $1
		RETURNING *;
	`
	args := []any{id, id_quest, id_team, name}

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
