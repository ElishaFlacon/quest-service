package models

type QuestTeamUser struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdUser      string `json:"idUser"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type QuestTeamUsers struct {
	IdQuestTeam int     `json:"idQuestTeam"`
	Users       []*User `json:"users"`
}
