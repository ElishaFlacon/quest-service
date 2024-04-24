package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/jackc/pgx/v5"
)

type TQuestTeamUser struct {
	table string
}

var QuestTeamUser = &TQuestTeamUser{
	table: "quest_team_user",
}

func (*TQuestTeamUser) Get(id int) ([]*models.QuestTeamUser, error) {
	sqlString := `
		SELECT * FROM "quest_team_user" 
		WHERE id_quest_team_user = $1;
	`

	data, err := database.BaseQuery[models.QuestTeamUser](sqlString, id)

	return data, err
}

func (*TQuestTeamUser) GetAll() ([]*models.QuestTeamUser, error) {
	sqlString := `SELECT * FROM "quest_team_user";`

	data, err := database.BaseQuery[models.QuestTeamUser](sqlString)

	return data, err
}

func (*TQuestTeamUser) GetByQuestId(
	id int,
) ([]*models.QuestTeamUser, error) {
	// TODO для Тимура, 1 очередь
	// нужно дописать sql запрос, который будет выдавать поля из models.QuestTeamUser по айдишнику
	sqlString := `
		SELECT 
		"quest_team_user".id_quest_team,
		"quest_team_user".id_user   
		FROM "quest"
		

		ТУТ ДОЛЖНА БЫТЬ ЦЕПОЧКА INNER JOIN'ов
		схема таблицы в readme (смотреть на гитхабе, чтобы картинка подгрузилась)

		
		WHERE "quest".quest_id = $1;
	`

	data, err := database.BaseQuery[models.QuestTeamUser](sqlString, id)

	return data, err
}

func (init *TQuestTeamUser) CreateWithCopy(
	rows [][]any,
) (int64, error) {
	columnNames := []string{"id_user"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (*TQuestTeamUser) CreateWithBatch(
	questTeams []*models.QuestTeamUsers,
) ([]*models.QuestTeam, error) {
	sqlString := `
		INSERT INTO "quest_team_user"
		(id_quest_team, id_user)
		VALUES ($1, $2)
		RETURNING *;
	`

	batch := &pgx.Batch{}

	for index := range questTeams {
		questTeam := questTeams[index]
		for userIndex := range questTeam.Users {
			user := questTeam.Users[userIndex]
			batch.Queue(
				sqlString,
				questTeam.IdQuestTeam,
				user,
			)
		}
	}

	data, err := database.SendBatch[models.QuestTeam](batch)

	return data, err
}

func (*TQuestTeamUser) Update(
	id int,
	id_user string,
) ([]*models.QuestTeamUser, error) {
	sqlString := `
		UPDATE "quest_team_user" 
		SET id_user = $2
		WHERE id_quest_team_user = $1
		RETURNING *;
	`
	args := []any{id, id_user}

	data, err := database.BaseQuery[models.QuestTeamUser](
		sqlString,
		args...,
	)

	return data, err
}

func (*TQuestTeamUser) Delete(id int) ([]*models.QuestTeamUser, error) {
	sqlString := `
		DELETE FROM "quest_team_user" 
		WHERE id_quest_team_user = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.QuestTeamUser](sqlString, id)

	return data, err
}

func (*TQuestTeamUser) DeleteAll() ([]*models.QuestTeamUser, error) {
	sqlString := `DELETE FROM "quest_team_user" RETURNING *;`

	data, err := database.BaseQuery[models.QuestTeamUser](sqlString)

	return data, err
}
