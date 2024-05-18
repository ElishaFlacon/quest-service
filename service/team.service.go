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

func (*TTeam) GetAllTeams(
	bearer string,
) ([]*models.IdeaServiceTeam, error) {
	client := &http.Client{}
	teams := &models.IdeaServiceTeams{}

	teamServiceUrl := utils.GetTeamServiceUrl()
	requestUrl := fmt.Sprintf("%s/all", teamServiceUrl)

	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)
	request.Header.Set("Authorization", bearer)

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
	token string,
	idTeams []string,
) ([]*models.IdeaServiceTeam, error) {
	allTeams, errAllTeams := Team.GetAllTeams(token)
	if errAllTeams != nil {
		return nil, errAllTeams
	}

	teams := []*models.IdeaServiceTeam{}

	// O(n^2) я плакал
	for _, team := range allTeams {
		for _, idTeam := range idTeams {
			if team.IdTeam == idTeam {
				teams = append(teams, team)
			}
		}
	}

	return teams, nil
}

func (*TTeam) GetUsersCount(
	token string,
	idTeams []string,
) (int, error) {
	teams, errTeams := Team.GetTeams(token, idTeams)
	if errTeams != nil {
		return 0, errTeams
	}

	count := 0

	for _, team := range teams {
		count += len(team.IdUsers)
	}

	return count, nil
}
