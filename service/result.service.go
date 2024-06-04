package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/jackc/pgx/v5"
)

type TResult struct {
	table string
}

var Result = &TResult{
	table: "result",
}

func (*TResult) Get(id int) ([]*models.Result, error) {
	sqlString := `SELECT * FROM "result" WHERE id_result = $1;`

	data, err := database.BaseQuery[models.Result](sqlString, id)

	return data, err
}

func (*TResult) GetByUserId(id string) ([]*models.Result, error) {
	sqlString := `SELECT * FROM "result" WHERE id_from_user = $1;`

	data, err := database.BaseQuery[models.Result](sqlString, id)

	return data, err
}

func (*TResult) GetByUsersId(ids []string) ([]*models.Result, error) {
	sqlString := `SELECT * FROM "result" WHERE id_result = $1;`

	batch := &pgx.Batch{}

	for _, id := range ids {
		batch.Queue(sqlString, id)
	}

	data, err := database.SendBatch[models.Result](batch)

	return data, err
}

func (*TResult) GetUsersCountByQuestId(id int) ([]*models.Count, error) {
	sqlString := `
		SELECT COUNT(*)
		FROM "result"
		WHERE id_quest = $1
		GROUP BY id_from_user;
	`

	data, err := database.BaseQuery[models.Count](sqlString, id)

	return data, err
}

func (*TResult) GetAll() ([]*models.Result, error) {
	sqlString := `SELECT * FROM "result";`

	data, err := database.BaseQuery[models.Result](sqlString)

	return data, err
}

func (init *TResult) CreateWithCopy(rows [][]any) (int64, error) {
	columnNames := []string{
		"id_quest",
		"id_indicator",
		"id_from_user",
		"id_to_user",
		"value",
	}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (*TResult) Update(
	id int,
	id_quest int,
	id_indicator int,
	id_from_user int,
	id_to_user int,
	value string,
) ([]*models.Result, error) {
	sqlString := `
		UPDATE "result" 
		SET (id_indicator, id_quest, id_from_user, id_to_user, value)
		VALUES ($2, $3, $4, $5, $6) 
		WHERE id_result = $1
		RETURNING *;
	`
	args := []any{id, id_indicator, id_quest, id_from_user, id_to_user, value}

	data, err := database.BaseQuery[models.Result](
		sqlString,
		args...,
	)

	return data, err
}

func (*TResult) Delete(id int) ([]*models.Result, error) {
	sqlString := `DELETE FROM "result" WHERE id_result = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Result](sqlString, id)

	return data, err
}

func (*TResult) DeleteAll() ([]*models.Result, error) {
	sqlString := `DELETE FROM "result" RETURNING *;`

	data, err := database.BaseQuery[models.Result](sqlString)

	return data, err
}
