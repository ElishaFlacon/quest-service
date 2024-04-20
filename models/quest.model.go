package models

type Quest struct {
	IdQuest     int    `json:"id"`
	IdTemplate  int    `json:"idTemplate"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	StartAt     int    `json:"startAt"`
	EndAt       int    `json:"endAt"`
}

type QuestResponse struct {
	IdQuest int     `json:"id"`
	Name    string  `json:"name"`
	StartAt int     `json:"startAt"`
	EndAt   int     `json:"endAt"`
	Percent float32 `json:"percent"`
	Status  string  `json:"status"`
}

type QuestWithIndicators struct {
	IdQuest    int         `json:"id"`
	Name       string      `json:"name"`
	StartAt    int         `json:"startAt"`
	EndAt      int         `json:"endAt"`
	Percent    float32     `json:"percent"`
	Status     string      `json:"status"`
	Indicators []Indicator `json:"indicators"`
}
