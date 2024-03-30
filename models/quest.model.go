package models

type Quest struct {
	IdQuest     int    `json:"idQuest"`
	Available   bool   `json:"available"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
