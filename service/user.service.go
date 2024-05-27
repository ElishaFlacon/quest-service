package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TUser struct{}

var User TUser

// Не работает, не реализована логика в другом микросервисе
func (*TUser) GetUsersInfo(
	IdUsers []*models.UsersFromAndToByResultId,
) ([]*models.UsersInfo, error) {
	client := &http.Client{}
	users := models.UsersInfoResponse{}

	jsonIdUsers, _ := json.Marshal(IdUsers)
	usersData := []byte(`{"usersFromAndToByResultId":` + string(jsonIdUsers))

	profileServiceUrl := utils.GetProfileServiceUrl()
	requestUrl := fmt.Sprintf("%s/users/all", profileServiceUrl)

	request, err := http.NewRequest(
		"GET",
		requestUrl,
		bytes.NewBuffer(usersData),
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &users)
	if err != nil {
		return nil, err
	}

	return users.UsersInfo, nil
}
