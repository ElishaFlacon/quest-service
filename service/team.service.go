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

func GetTeam(
	bearer string,
	id string,
) (*models.Team, error) {
	client := &http.Client{}
	team := &models.Team{}

	teamServiceUrl := utils.GetTeamServiceUrl()
	requestUrl := fmt.Sprintf("%s/%s", teamServiceUrl, id)

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

	err = json.Unmarshal(body, team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (*TTeam) GetTeams(
	bearer string,
	ids []string,
) ([]*models.Team, error) {
	teams := []*models.Team{}

	for _, id := range ids {
		team, errTeam := GetTeam(bearer, id)
		if errTeam != nil {
			teams = append(teams, nil)
			continue
		}

		teams = append(teams, team)
	}

	return teams, nil
}
