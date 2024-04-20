package models

type QuestTeamUser struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdUser      string `json:"idUser"`
}

type QuestTeamUsers struct {
	IdQuestTeam int      `json:"idQuestTeam"`
	Users       []string `json:"idUsers"`
}
