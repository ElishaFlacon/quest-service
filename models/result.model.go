package models

type Result struct {
	IdResult    int    `json:"id"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  int    `json:"idFromUser"`
	IdToUser    int    `json:"idToUser"`
	Value       string `json:"value"`
}
