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

func (*TTeam) GetTeam(
	bearer string,
	id string,
) (*models.Team, error) {
	client := &http.Client{}
	team := &models.TeamWithFullUser{}

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

	users := []*models.User{}
	for _, user := range team.Users {
		name := fmt.Sprintf(
			"%s %s",
			user.FirstName,
			user.LastName,
		)
		element := &models.User{
			IdUser: user.IdUser,
			Email:  user.Email,
			Name:   name,
		}

		users = append(users, element)
	}

	resultTeam := &models.Team{
		IdTeam: team.IdTeam,
		Name:   team.Name,
		Users:  users,
	}

	return resultTeam, nil
}

func (*TTeam) GetTeams(
	bearer string,
	ids []string,
) ([]*models.Team, error) {
	teams := []*models.Team{}

	for _, id := range ids {
		team, errTeam := Team.GetTeam(bearer, id)
		if errTeam != nil {
			return nil, errTeam
		}

		teams = append(teams, team)
	}

	return teams, nil
}
