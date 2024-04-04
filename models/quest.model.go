package models

type Quest struct {
	IdQuest     int    `json:"idQuest"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

// для списка опросов
type QuestListResponse struct {
	IdQuest int    `json:"idQuest"`
	Name    string `json:"name"`
	StartAt int    `json:"startAt"`
	EndAt   int    `json:"endAt"`
	Percent string `json:"percent"`
	Status  string `json:"status"`
}
