package models

type QuestTeam struct {
	IdQuestTeam int    `json:"id"`
	IdTeam          string `json:"idTeam"`
	IdQuest         string `json:"idQuest"`
}
