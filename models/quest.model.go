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
	IdQuest   int     `json:"id"`
	Name      string  `json:"name"`
	StartAt   string  `json:"startAt"`
	EndAt     string  `json:"endAt"`
	Percent   float32 `json:"percent"`
	Status    string  `json:"status"`
	Available bool    `json:"available"`
}

type QuestWithIndicators struct {
	IdQuest    int                          `json:"id"`
	Name       string                       `json:"name"`
	StartAt    string                       `json:"startAt"`
	EndAt      string                       `json:"endAt"`
	Percent    float32                      `json:"percent"`
	Status     string                       `json:"status"`
	Available  bool                         `json:"available"`
	Indicators []*IndicatorWithCategoryName `json:"indicators"`
}

type QuestWithUsers struct {
	IdQuest   int      `json:"id"`
	Name      string   `json:"name"`
	StartAt   string   `json:"startAt"`
	EndAt     string   `json:"endAt"`
	Percent   float32  `json:"percent"`
	Status    string   `json:"status"`
	Available bool     `json:"available"`
	Users     []string `json:"users"`
}

type QuestWithUsersAndIndicators struct {
	IdQuest    int                          `json:"id"`
	Name       string                       `json:"name"`
	StartAt    string                       `json:"startAt"`
	EndAt      string                       `json:"endAt"`
	Percent    float32                      `json:"percent"`
	Status     string                       `json:"status"`
	Available  bool                         `json:"available"`
	Indicators []*IndicatorWithCategoryName `json:"indicators"`
	Users      []string                     `json:"users"`
}

type QuestCreateRequest struct {
	IdTemplate  int          `json:"idTemplate"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Teams       []Id[string] `json:"teams"`
	StartAt     int          `json:"startAt"`
	EndAt       int          `json:"endAt"`
}
