package models

type Quest struct {
	IdQuest     int    `json:"idQuest"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}
