package models

type Quest struct {
	IdQuest     int    `json:"idQuest"`
	IdTemplate  int    `json:"idQuestTemplate"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	StartAt     int    `json:"startAt"`
	EndAt       int    `json:"endAt"`
}

type QuestWithPercent struct {
	IdQuest     int    `json:"idQuest"`
	IdTemplate  int    `json:"idQuestTemplate"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Percent     int    `json:"progress"`
	Available   bool   `json:"available"`
	StartAt     int    `json:"startAt"`
	EndAt       int    `json:"endAt"`
}

type QuestResponse struct {
	IdQuest    int    `json:"idQuest"`
	IdTemplate int    `json:"idQuestTemplate"`
	Name       string `json:"name"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
	Percent    int    `json:"progress"`
	Status     string `json:"status"`
	Available  bool   `json:"available"`
}

type QuestWithStatuses struct {
	IdQuest    int               `json:"idQuest"`
	IdTemplate int               `json:"idQuestTemplate"`
	Name       string            `json:"name"`
	StartAt    string            `json:"startAt"`
	EndAt      string            `json:"endAt"`
	Percent    int               `json:"progress"`
	Status     string            `json:"status"`
	Available  bool              `json:"available"`
	Teams      []*TeamWithStatus `json:"teams"`
}

type QuestWithStatusesForUser struct {
	IdQuest    int               `json:"idQuest"`
	IdTemplate int               `json:"idQuestTemplate"`
	Name       string            `json:"name"`
	StartAt    string            `json:"startAt"`
	EndAt      string            `json:"endAt"`
	Percent    int               `json:"progress"`
	Status     string            `json:"status"`
	IsPass     bool              `json:"isPass"`
	Available  bool              `json:"available"`
	Teams      []*TeamWithStatus `json:"teams"`
}

type QuestWithIndicators struct {
	IdQuest    int                          `json:"idQuest"`
	IdTemplate int                          `json:"idQuestTemplate"`
	Name       string                       `json:"name"`
	StartAt    string                       `json:"startAt"`
	EndAt      string                       `json:"endAt"`
	Percent    int                          `json:"progress"`
	Status     string                       `json:"status"`
	Available  bool                         `json:"available"`
	Indicators []*IndicatorWithCategoryName `json:"indicators"`
}

type QuestWithUsers struct {
	IdQuest    int      `json:"idQuest"`
	IdTemplate int      `json:"idQuestTemplate"`
	Name       string   `json:"name"`
	StartAt    string   `json:"startAt"`
	EndAt      string   `json:"endAt"`
	Percent    int      `json:"progress"`
	Status     string   `json:"status"`
	Available  bool     `json:"available"`
	Users      []string `json:"users"`
}

type QuestWithUsersAndIndicators struct {
	IdQuest    int                          `json:"idQuest"`
	IdTemplate int                          `json:"idQuestTemplate"`
	Name       string                       `json:"name"`
	StartAt    string                       `json:"startAt"`
	EndAt      string                       `json:"endAt"`
	Percent    int                          `json:"progress"`
	Status     string                       `json:"status"`
	Available  bool                         `json:"available"`
	Indicators []*IndicatorWithCategoryName `json:"indicators"`
	Users      []string                     `json:"users"`
}

type QuestCreateRequest struct {
	IdTemplate  int          `json:"idQuestTemplate"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Teams       []Id[string] `json:"teams"`
	StartAt     int          `json:"startAt"`
	EndAt       int          `json:"endAt"`
}
