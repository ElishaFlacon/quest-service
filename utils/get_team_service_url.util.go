package utils

import (
	"os"
)

func GetTeamServiceUrl() string {
	return os.Getenv("TEAM_SERVICE_URL")
}
