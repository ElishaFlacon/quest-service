package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type ILaunchQuest interface {
	GetLaunchQuest(id int) ([]*models.LaunchQuest, error)
	GetAllLaunchQuests() ([]*models.LaunchQuest, error)
	CreateLaunchQuest(launchQuest models.LaunchQuest) ([]*models.LaunchQuest, error)
	UpdateLaunchQuest(rows [][]any) (int64, error)
	DeleteLaunchQuest(id int) ([]*models.LaunchQuest, error)
	DeleteAllLaunchQuests() ([]*models.LaunchQuest, error)
}

var LaunchQuest ILaunchQuest

func (repo *TDatabase) GetLaunchQuest(id int) ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM launch_quests WHERE id_launch_quest = $1;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllLaunchQuests() ([]*models.LaunchQuest, error) {
	sqlString := `SELECT * FROM launch_quests;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query)

	return data, err
}

func (repo *TDatabase) CreateLaunchQuest(launchQuest models.LaunchQuest) ([]*models.LaunchQuest, error) {
	sqlString := `INSERT INTO launch_quest (id_quest, id_team, available, start_at, end_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](
		sqlString,
		repo.db.Query,
		launchQuest.IdQuest,
		launchQuest.IdTeam,
		launchQuest.Available,
		launchQuest.StartAt,
		launchQuest.EndAt,
	)
	return data, err
}

func (repo *TDatabase) UpdateLaunchQuest(rows [][]any) (int64, error) {
	tableName := "launch_quest"
	columnNames := []string{"available", "end_at"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) DeleteLaunchQuest(id int) ([]*models.LaunchQuest, error) {
	sqlString := `DELETE FROM launch_quest WHERE id_launch_quest = $1 RETURNING *;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query, id)
	return data, err
}

func (repo *TDatabase) DeleteAllLaunchQuests() ([]*models.LaunchQuest, error) {
	sqlString := `DELETE FROM launch_quest RETURNING *;`

	data, err := core.QueryWithReturningData[models.LaunchQuest](sqlString, repo.db.Query)
	return data, err
}
