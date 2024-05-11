package models

type Team struct {
	IdTeam string `json:"idTeam"`
	Users  []User `json:"users"`
}

type IdeaServiceTeam struct {
	IdTeam  string
	IdUsers []string
}

type IdeaServiceTeams struct {
	Teams []*IdeaServiceTeam `json:"teams"`
}
