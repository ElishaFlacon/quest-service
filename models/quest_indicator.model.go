package models

type QuestIndicator struct {
	IdQuestIndicator int `json:"id_quest_indicator"`
	IdQuest          int `json:"id_quest"`
	IdIndicator      int `json:"id_indicator"`
}
