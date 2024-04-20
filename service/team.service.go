package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TTeam struct{}

var Team *TTeam

var TeamServiceUrl = utils.GetTeamServiceUrl()

func (*TTeam) GetAllMembersOfTeams(
	idTeams []*int,
) ([]*models.TeamMembers, error) {
	client := &http.Client{}
	teams := models.TeamsResponse{}

	jsonIdTeams, _ := json.Marshal(idTeams)
	teamsData := []byte(`{"idTeams":` + string(jsonIdTeams))

	request, err := http.NewRequest(
		"GET",
		TeamServiceUrl,
		bytes.NewBuffer(teamsData),
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

	err = json.Unmarshal(responseBody, &teams)
	if err != nil {
		return nil, err
	}

	return teams.TeamMembers, nil

}
