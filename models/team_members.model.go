package models

type TeamMembers struct {
	IdTeam  string
	IdUsers []string
}

type TeamsResponse struct {
	TeamMembers []*TeamMembers `json:"teams"`
}
