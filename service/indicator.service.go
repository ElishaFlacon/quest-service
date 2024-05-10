package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TIndicator struct {
	table string
}

var Indicator = &TIndicator{
	table: "indicator",
}

func (*TIndicator) Get(id int) (*models.IndicatorWithCategoryName, error) {
	sqlString := `
	SELECT
	"indicator".id_indicator, 
	"indicator".name, 
	"indicator".description,
	"indicator".id_category,
	"indicator".visible,
	CASE 
		WHEN 
			"indicator".from_role = '' 
				OR 
			"indicator".from_role IS NULL 
				THEN 'ALL' 
		ELSE "indicator".from_role END, 
	CASE 
		WHEN 
			"indicator".to_role = '' 
				OR 
			"indicator".to_role IS NULL 
				THEN 'ALL' 
		ELSE "indicator".to_role END, 
	"category".name as category_name
	FROM "indicator"
	INNER JOIN "category" ON 
		"indicator".id_category = "category".id_category
	WHERE "indicator".id_indicator = $1;
`

	data, err := database.BaseQuery[models.IndicatorWithCategoryName](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TIndicator) GetByTemplateId(id int) ([]*models.IndicatorWithCategoryName, error) {
	sqlString := `
		SELECT 
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".id_category,
			"indicator".visible,
			CASE 
				WHEN 
					"indicator".from_role = '' 
						OR 
					"indicator".from_role IS NULL 
						THEN 'ALL' 
				ELSE "indicator".from_role END, 
			CASE 
				WHEN 
					"indicator".to_role = '' 
						OR 
					"indicator".to_role IS NULL 
						THEN 'ALL' 
				ELSE "indicator".to_role END, 
			"category".name as category_name
		FROM "indicator"
		INNER JOIN "category" 
			ON "indicator".id_category = "category".id_category
		INNER JOIN "template_indicator" 
			ON "indicator".id_indicator = "template_indicator".id_indicator
		INNER JOIN "template" 
			ON "template_indicator".id_template = "template".id_template
		WHERE "template".id_template = $1;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategoryName](sqlString, id)

	return data, err
}

func (*TIndicator) GetByQuestId(id int) ([]*models.IndicatorWithCategoryName, error) {
	sqlString := `
		SELECT 
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".id_category,
			"indicator".visible,
			CASE 
				WHEN 
					"indicator".from_role = '' 
						OR 
					"indicator".from_role IS NULL 
						THEN 'ALL' 
				ELSE "indicator".from_role END, 
			CASE 
				WHEN 
					"indicator".to_role = '' 
						OR 
					"indicator".to_role IS NULL 
						THEN 'ALL' 
				ELSE "indicator".to_role END, 
			"category".name as category_name
		FROM "indicator"
		INNER JOIN "category" 
			ON "indicator".id_category = "category".id_category
		INNER JOIN "template_indicator" ON 
			"indicator".id_indicator = "template_indicator".id_indicator
		INNER JOIN "template" ON 
			"template_indicator".id_template = "template".id_template
		INNER JOIN "quest" ON 
			"template".id_template = "quest".id_template
		WHERE "quest".id_quest = $1;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategoryName](sqlString, id)

	return data, err
}

func (*TIndicator) GetAll() ([]*models.IndicatorWithCategoryName, error) {
	sqlString := `
		SELECT
		"indicator".id_indicator, 
		"indicator".name, 
		"indicator".description,
		"indicator".id_category,
		"indicator".visible,
		CASE 
			WHEN 
				"indicator".from_role = '' 
					OR 
				"indicator".from_role IS NULL 
					THEN 'ALL' 
			ELSE "indicator".from_role END, 
		CASE 
			WHEN 
				"indicator".to_role = '' 
					OR 
				"indicator".to_role IS NULL 
					THEN 'ALL' 
			ELSE "indicator".to_role END, 
		"category".name as category_name
		FROM "indicator"
		INNER JOIN "category" ON 
			"indicator".id_category = "category".id_category;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategoryName](sqlString)

	return data, err
}

func (*TIndicator) Create(
	idCategory int,
	name string,
	description string,
	fromRole string,
	toRole string,
) (*models.Indicator, error) {
	sqlString := `
		INSERT INTO "indicator" 
		(name, description, from_role, to_role, id_category) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *;
	`
	args := []any{name, description, fromRole, toRole, idCategory}

	data, err := database.BaseQuery[models.Indicator](sqlString, args...)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TIndicator) Hide(id int) (*models.Indicator, error) {
	sqlString := `
		UPDATE "indicator" 
		SET visible = false
		WHERE id_indicator = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TIndicator) Delete(id int) (*models.Indicator, error) {
	sqlString := `
		DELETE FROM "indicator" 
		WHERE id_indicator = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}
