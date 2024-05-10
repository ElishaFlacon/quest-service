package models

type QuestTeam struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdTeam      string `json:"idTeam"`
	IdQuest     string `json:"idQuest"`
}
