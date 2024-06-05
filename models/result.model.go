package models

type Result struct {
	IdResult    int    `json:"idResult"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  string `json:"idFromUser"`
	IdToUser    string `json:"idToUser"`
	Value       string `json:"value"`
}

type ResultResponse struct {
	CreatedResults int64 `json:"createdResults"`
}

type ResultCreate struct {
	Results []struct {
		IdQuest     int    `json:"idQuest"`
		IdIndicator int    `json:"idIndicator"`
		IdFromUser  string `json:"idFromUser"`
		IdToUser    string `json:"idToUser"`
		Value       string `json:"value"`
	} `json:"results"`
}
