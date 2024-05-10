package models

type Team struct {
	IdTeam string `json:"idTeam"`
	Users  []User `json:"users"`
}

type TeamMembers struct {
	IdTeam  string
	IdUsers []string
}

type TeamsResponse struct {
	TeamMembers []*TeamMembers `json:"teams"`
}
