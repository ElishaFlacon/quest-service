package service

import (
	"bytes"
	"encoding/csv"
	"errors"

	"github.com/ElishaFlacon/quest-service/models"
)

type TCsv struct{}

var Csv = &TCsv{}

func findUsersInfoByIdResult(
	idResult int,
	usersInfo []*models.UsersInfo,
) (*models.UsersInfo, error) {
	for _, user := range usersInfo {
		if user.IdResult == idResult {
			return user, nil
		}
	}
	return nil, errors.New("ошибка")
}

func (*TCsv) GetCsvTableByQuestId(
	bearer string,
	IdQuest int,
) ([]byte, error) {
	quest, err := Quest.Get(bearer, IdQuest)
	if err != nil {
		return nil, err
	}

	results, errResults := Result.GetByQuestId(IdQuest)
	if errResults != nil {
		return nil, errResults
	}

	var usersFromAndToByResultId []*models.UsersFromAndToByResultId

	var csvBuffer bytes.Buffer

	for _, result := range results {
		userFromAndToByResultId := &models.UsersFromAndToByResultId{}
		userFromAndToByResultId.IdToUser = result.IdToUser
		userFromAndToByResultId.IdFromUser = result.IdFromUser
		userFromAndToByResultId.IdResult = result.IdResult
		usersFromAndToByResultId = append(
			usersFromAndToByResultId,
			userFromAndToByResultId,
		)
	}

	usersInfo, errUsersInfo := User.GetToAndFromUsers(usersFromAndToByResultId)
	if errUsersInfo != nil {
		return nil, errUsersInfo
	}

	writer := csv.NewWriter(&csvBuffer)
	defer writer.Flush()

	columns := []string{
		"Название опроса",
		"Название вопроса",
		"От кого",
		"Роль (от кого)",
		"Кому",
		"Роль (кому)",
		"Результат ответа",
	}

	errWriter := writer.Write(columns)
	if errWriter != nil {
		return nil, errWriter
	}

	for _, result := range results {
		indicator, errIndicator := Indicator.Get(result.IdIndicator)
		if errIndicator != nil {
			return nil, errIndicator
		}

		userInfo, errUserInfo := findUsersInfoByIdResult(result.IdResult, usersInfo)
		if errUserInfo != nil {
			return nil, errUserInfo
		}
		err = writer.Write([]string{
			quest.Name,
			indicator.Name,
			userInfo.FullNameFromUser,
			indicator.FromRole,
			userInfo.FullNameToUser,
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
