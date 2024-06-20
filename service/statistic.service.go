package service

import (
	"bytes"
	"encoding/csv"

	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
)

type TStatistic struct{}

var Statistic = &TStatistic{}

func (*TStatistic) GetQuestStatistic(id int) ([]byte, error) {
	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)

	columns := []string{
		"Название опроса",
		"Название вопроса",
		"Название команды",
		"От кого",
		"Роль (от кого)",
		"Кому",
		"Роль (кому)",
		"Результат ответа",
	}

	if err := writer.Write(columns); err != nil {
		return nil, err
	}

	sqlString := `
		SELECT 
			q.name AS quest_name,
			i.name AS indicator_name,
			u1.name AS from_user,
			i.from_role,
			u2.name AS to_user,
			i.to_role,
			r.value,
			qt.name AS team_name
		FROM result r
		INNER JOIN quest q ON r.id_quest = q.id_quest
		INNER JOIN indicator i ON r.id_indicator = i.id_indicator
		INNER JOIN quest_team_user u1 ON r.id_from_user = u1.id_user
		INNER JOIN quest_team_user u2 ON r.id_to_user = u2.id_user
		INNER JOIN quest_team qt ON u1.id_quest_team = qt.id_quest_team
		WHERE r.id_quest = $1;
	`

	data, err := database.BaseQuery[models.StatisticQuestCsvRecord](sqlString, id)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		if err := writer.Write(record.ToSlice()); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return csvBuffer.Bytes(), nil
}
