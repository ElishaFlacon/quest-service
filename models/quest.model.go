package models

type Quest struct {
	IdQuest     int    `json:"idQuest"`
	IdTemplate  int    `json:"idTemplate"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	StartAt     int    `json:"startAt"`
	EndAt       int    `json:"endAt"`
}

// для списка опросов
type QuestListResponse struct {
	IdQuest int     `json:"idQuest"`
	Name    string  `json:"name"`
	StartAt int     `json:"startAt"`
	EndAt   int     `json:"endAt"`
	Percent float32 `json:"percent"`
	Status  string  `json:"status"`
}
