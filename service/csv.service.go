package service

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/ElishaFlacon/quest-service/models"
)

type TCsv struct{}

var Csv *TCsv

func findUsersInfoByIdResult(idResult int, usersInfo []*models.UsersInfo) (*models.UsersInfo, error) {
	for _, user := range usersInfo {
		if user.IdResult == idResult {
			return user, nil
		}
	}
	return nil, errors.New("ошибка")
}

func (*TCsv) GetCsvTableByQuestId(IdQuest int) ([]byte, error) {
	quest, err := Quest.Get(IdQuest)
	if err != nil {
		return nil, err
	}

	results, err := Result.GetByQuestId(IdQuest)

	var usersFromAndToByResultId []*models.UsersFromAndToByResultId

	var csvBuffer bytes.Buffer

	for _, result := range results {
		var userFromAndToByResultId *models.UsersFromAndToByResultId
		userFromAndToByResultId.IdToUser = result.IdToUser
		userFromAndToByResultId.IdFromUser = result.IdFromUser
		userFromAndToByResultId.IdResult = result.IdResult
		_ = append(usersFromAndToByResultId, userFromAndToByResultId)
	}

	usersInfo, err := User.GetToAndFromUsers(usersFromAndToByResultId)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(&csvBuffer)
	defer writer.Flush()

	err = writer.Write([]string{"Название опроса", "Название вопроса",
		"От кого", "Роль (от кого)", "Кому", "Роль (кому)", "Результат ответа"})
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		indicator, err := Indicator.Get(result.IdIndicator)
		if err != nil {
			return nil, err
		}

		userInfoByIdResult, err := findUsersInfoByIdResult(result.IdResult, usersInfo)
		if err != nil {
			return nil, err
		}
		err = writer.Write([]string{
			quest.Name,
			indicator.Name,
			userInfoByIdResult.FullNameFromUser,
			indicator.FromRole,
			userInfoByIdResult.FullNameToUser,
			indicator.ToRole,
			result.Value,
		})
		if err != nil {
			return nil, err
		}
	}

	csvContent := csvBuffer.Bytes()

	return csvContent, nil
}
