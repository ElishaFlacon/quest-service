package models

type QuestTeam struct {
	IdQuestTeamUser int    `json:"id"`
	IdTeam          string `json:"idTeam"`
	IdQuest         string `json:"idQuest"`
}
