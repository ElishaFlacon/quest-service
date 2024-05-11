package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TTeam struct{}

var Team *TTeam

var TeamServiceUrl = utils.GetTeamServiceUrl()

func (*TTeam) GetAllTeams() ([]*models.IdeaServiceTeam, error) {
	client := &http.Client{}
	teams := &models.IdeaServiceTeams{}

	requestUrl := fmt.Sprintf("%s/all", TeamServiceUrl)

	request, err := http.NewRequest("GET", requestUrl, nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &teams)
	if err != nil {
		return nil, err
	}

	return teams.Teams, nil
}

func (*TTeam) GetTeams(
	idTeams []string,
) ([]*models.IdeaServiceTeam, error) {
	allTeams, errAllTeams := Team.GetAllTeams()
	if errAllTeams != nil {
		return nil, errAllTeams
	}

	teams := []*models.IdeaServiceTeam{}

	for _, team := range allTeams {
		for _, idTeam := range idTeams {
			if team.IdTeam == idTeam {
				teams = append(teams, team)
			}
		}
	}

	return teams, nil
}
