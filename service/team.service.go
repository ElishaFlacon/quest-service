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
) (*models.TeamWithStringUsers, error) {
	client := &http.Client{}
	team := models.Team{}

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

	err = json.Unmarshal(body, &team)
	if err != nil {
		return nil, err
	}

	users := []string{}
	for _, user := range team.Users {
		users = append(users, user.IdUser)
	}

	returnTeam := &models.TeamWithStringUsers{
		IdTeam: team.IdTeam,
		Users:  users,
	}

	return returnTeam, nil
}

func (*TTeam) GetAllTeams(
	bearer string,
) ([]models.Id[string], error) {
	client := &http.Client{}
	teams := []models.Id[string]{}

	teamServiceUrl := utils.GetTeamServiceUrl()
	requestUrl := fmt.Sprintf("%s/all", teamServiceUrl)

	request, errRequest := http.NewRequest("GET", requestUrl, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	request.Header.Set(
		"Accept",
		"application/json, text/plain, */*",
	)
	request.Header.Set(
		"Accept-Encoding",
		"gzip, deflate, br, zstd",
	)
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	)
	request.Header.Set("Authorization", bearer)

	response, errResponse := client.Do(request)
	if errResponse != nil {
		return nil, errResponse
	}
	defer response.Body.Close()

	body, errBody := io.ReadAll(response.Body)
	if errBody != nil {
		return nil, errBody
	}

	errUnmarshal := json.Unmarshal(body, &teams)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return teams, nil
}

func (*TTeam) GetTeams(
	bearer string,
	idTeams []string,
) ([]*models.TeamWithStringUsers, error) {
	allTeams, errAllTeams := Team.GetAllTeams(bearer)
	if errAllTeams != nil {
		return nil, errAllTeams
	}

	teams := []*models.TeamWithStringUsers{}

	// O(n^2) я плакал
	for _, team := range allTeams {
		for _, idTeam := range idTeams {
			if team.Id != idTeam {
				continue
			}

			teamRequest, errTeamRequest := GetTeam(bearer, team.Id)
			if errTeamRequest != nil {
				return nil, errTeamRequest
			}

			teams = append(teams, teamRequest)
		}
	}

	return teams, nil
}

func (*TTeam) GetUsersCount(
	bearer string,
	idTeams []string,
) (int, error) {
	teams, errTeams := Team.GetTeams(bearer, idTeams)
	if errTeams != nil {
		return 0, errTeams
	}

	count := 0

	for _, team := range teams {
		count += len(team.Users)
	}

	return count, nil
}
