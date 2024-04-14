package service

import (
	"bytes"
	"encoding/json"
	"github.com/ElishaFlacon/quest-service/models"
	"io"
	"log"
	"net/http"
)

type TTeam struct{}

var Team *TTeam

func (*TTeam) GetAllMembersOfTeams(idTeams []*int) ([]*models.TeamMembers, error) {
	var teams models.TeamsResponse

	apiUrl := "/api/v1/teamservice/get-all-members"

	jsonIdTeams, _ := json.Marshal(idTeams)
	teamsData := []byte(`{"idTeams":` + string(jsonIdTeams))

	request, err := http.NewRequest("GET", apiUrl, bytes.NewBuffer(teamsData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	err = json.Unmarshal(responseBody, &teams)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return teams.TeamMembers, nil

}
