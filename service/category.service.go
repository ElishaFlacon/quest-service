package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TCategory struct {
	table string
}

var Category = &TCategory{
	table: "category",
}

func (*TCategory) Get(id int) (*models.Category, error) {
	sqlString := `
		SELECT * FROM "category" 
		WHERE id_category = $1;
	`

	data, err := database.BaseQuery[models.Category](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TCategory) GetAll() ([]*models.Category, error) {
	sqlString := `
		SELECT * 
		FROM "category"
		WHERE available = true;
	`

	data, err := database.BaseQuery[models.Category](sqlString)

	return data, err
}

func (*TCategory) Create(name string) (*models.Category, error) {
	sqlString := `
		INSERT INTO "category" 
		(name) VALUES ($1) 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Category](sqlString, name)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TCategory) Hide(id int) (*models.Category, error) {
	sqlString := `
		UPDATE "category" 
		SET available = false
		WHERE id_category = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Category](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}
