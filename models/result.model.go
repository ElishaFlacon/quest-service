package models

type Result struct {
	IdResult      int    `json:"idResult"`
	IdIndicator   int    `json:"idIndicator"`
	IdLaunchQuest int    `json:"idLaunchQuest"`
	IdFromUser    int    `json:"idFromUser"`
	IdToUser      int    `json:"idToUser"`
	Value         string `json:"value"`
}
