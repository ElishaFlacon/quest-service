package models

type Result struct {
	IdResult    int    `json:"id"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  string `json:"idFromUser"`
	IdToUser    string `json:"idToUser"`
	Value       string `json:"value"`
}
