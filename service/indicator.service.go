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

func (*TIndicator) Get(id int) (*models.IndicatorWithCategory, error) {
	sqlString := `
		SELECT
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".answers,
			"indicator".id_category,
			"indicator".available,
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
		WHERE "indicator".id_indicator = $1 AND "indicator".available = true;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategory](sqlString, id)

	return utils.CultivateFirstDataElement(data, err)
}

func (*TIndicator) GetByTemplateId(id int) ([]*models.IndicatorWithCategory, error) {
	sqlString := `
		SELECT 
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".answers,
			"indicator".id_category,
			"indicator".available,
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
		WHERE "template".id_template = $1 AND "indicator".available = true;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategory](sqlString, id)

	return data, err
}

// NOTE: видно при available = false
func (*TIndicator) GetByQuestId(id int) ([]*models.IndicatorWithCategory, error) {
	sqlString := `
		SELECT 
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".answers,
			"indicator".id_category,
			"indicator".available,
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

	data, err := database.BaseQuery[models.IndicatorWithCategory](sqlString, id)

	return data, err
}

func (*TIndicator) GetAll() ([]*models.IndicatorWithCategory, error) {
	sqlString := `
		SELECT
			"indicator".id_indicator, 
			"indicator".name, 
			"indicator".description,
			"indicator".answers,
			"indicator".id_category,
			"indicator".available,
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
		WHERE "indicator".available = true;
	`

	data, err := database.BaseQuery[models.IndicatorWithCategory](sqlString)

	return data, err
}

func (*TIndicator) Create(
	idCategory int,
	name string,
	description string,
	answers []string,
	fromRole string,
	toRole string,
) (*models.Indicator, error) {
	sqlString := `
		INSERT INTO "indicator" 
		(name, description, answers, from_role, to_role, id_category) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING *;
	`
	args := []any{name, description, answers, fromRole, toRole, idCategory}

	data, err := database.BaseQuery[models.Indicator](sqlString, args...)

	return utils.CultivateFirstDataElement(data, err)
}

func (*TIndicator) Hide(id int) (*models.Indicator, error) {
	sqlString := `
		UPDATE "indicator" 
		SET available = false
		WHERE id_indicator = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Indicator](sqlString, id)

	return utils.CultivateFirstDataElement(data, err)
}
