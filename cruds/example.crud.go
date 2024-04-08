package cruds

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TExample struct {
	table string
}

var Example = &TExample{
	table: "example",
}

func (init *TExample) Get(id int) ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example" WHERE id=$1;`

	data, err := database.BaseQuery[models.Example](sqlString, id)

	return data, err
}

func (init *TExample) GetAll() ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example";`

	data, err := database.BaseQuery[models.Example](sqlString)

	return data, err
}

func (init *TExample) Create(rows [][]any) (int64, error) {
	columnNames := []string{"id", "value"}

	count, err := database.CopyFromQuery(init.table, columnNames, rows)

	return count, err
}

func (init *TExample) Update(id int, value string) ([]*models.Example, error) {
	sqlString := `UPDATE "example" SET value=$2 WHERE id=$1 RETURNING *;`
	args := []any{id, value}

	data, err := database.BaseQuery[models.Example](sqlString, args...)

	return data, err
}

func (init *TExample) Delete(id int) ([]*models.Example, error) {
	sqlString := `DELETE * FROM "example" WHERE id_indicator = $1 RETURNING *;`

	data, err := database.BaseQuery[models.Example](sqlString, id)

	return data, err
}

func (init *TExample) DeleteAll() ([]*models.Example, error) {
	sqlString := `DELETE * FROM "example" RETURNING *;`

	data, err := database.BaseQuery[models.Example](sqlString)

	return data, err
}
