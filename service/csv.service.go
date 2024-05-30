package service

import (
	"bytes"
	"encoding/csv"
)

type TCsv struct{}

var Csv = &TCsv{}

func (*TCsv) GetCsvByQuestId(bearer string, id int) ([]byte, error) {
	_, err := Quest.Get(bearer, id)
	if err != nil {
		return nil, err
	}

	results, errResults := Result.GetByQuestId(id)
	if errResults != nil {
		return nil, errResults
	}

	// var usersFromAndToByResultId []*models.UsersFromAndToByResultId

	var csvBuffer bytes.Buffer

	// for _, result := range results {
	// 	userFromAndToByResultId := &models.UsersFromAndToByResultId{}
	// 	userFromAndToByResultId.IdToUser = result.IdToUser
	// 	userFromAndToByResultId.IdFromUser = result.IdFromUser
	// 	userFromAndToByResultId.IdResult = result.IdResult
	// 	usersFromAndToByResultId = append(
	// 		usersFromAndToByResultId,
	// 		userFromAndToByResultId,
	// 	)
	// }

	// в этом месте проблема, метод на другом микросервисе не был реализован
	// usersInfo, errUsersInfo := User.GetUsersInfo(usersFromAndToByResultId)
	// if errUsersInfo != nil {
	// 	return nil, errUsersInfo
	// }

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
		_, errIndicator := Indicator.Get(result.IdIndicator)
		if errIndicator != nil {
			return nil, errIndicator
		}

		// userInfo, errUserInfo := findUsersInfoByIdResult(
		// 	result.IdResult,
		// 	usersInfo,
		// )
		// if errUserInfo != nil {
		// 	return nil, errUserInfo
		// }
		// err = writer.Write([]string{
		// 	quest.Name,
		// 	indicator.Name,
		// 	userInfo.FullNameFromUser,
		// 	indicator.FromRole,
		// 	userInfo.FullNameToUser,
		// 	indicator.ToRole,
		// 	result.Value,
		// })
		// if err != nil {
		// 	return nil, err
		// }
	}

	csvContent := csvBuffer.Bytes()

	return csvContent, nil
}
