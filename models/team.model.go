package models

type Member struct {
	IdUser string `json:"id"`
	Email  string `json:"email"`
}

type Team struct {
	IdTeam string    `json:"id"`
	Users  []*Member `json:"members"`
}

type TeamWithStringUsers struct {
	IdTeam string   `json:"id"`
	Users  []string `json:"users"`
}
