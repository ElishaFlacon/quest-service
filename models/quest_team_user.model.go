package models

type QuestTeamUser struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdUser      string `json:"idUser"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type QuestTeamUserWithTeam struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdTeam      int    `json:"idTeam"`
	IdUser      string `json:"idUser"`
}

type QuestTeamUsers struct {
	IdQuestTeam int     `json:"idQuestTeam"`
	Users       []*User `json:"users"`
}

type QuestTeamUsersWithFullUsers struct {
	IdQuestTeam int `json:"idQuestTeam"`
	Users       []struct {
		IdUser string `json:"id"`
		Name   int    `json:"name"`
		IsPass bool   `json:"progress"`
	} `json:"users"`
}
