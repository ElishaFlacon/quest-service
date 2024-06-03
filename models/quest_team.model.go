package models

type QuestTeam struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdQuest     int    `json:"idQuest"`
	IdTeam      string `json:"idTeam"`
	Name        string `json:"name"`
}

type QuestTeamWithPercent struct {
	IdQuestTeam int    `json:"idQuestTeam"`
	IdQuest     int    `json:"idQuest"`
	IdTeam      string `json:"idTeam"`
	Name        string `json:"name"`
	Percent     int    `json:"percent"`
}
