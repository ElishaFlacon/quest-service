package models

type LaunchQuest struct {
	IdLaunchQuest int  `json:"idLaunchQuest"`
	IdQuest       int  `json:"idQuest"`
	IdTeam        int  `json:"idTeam"`
	Available     bool `json:"available"`
	StartAt       int  `json:"startAt"`
	EndAt         int  `json:"endAt"`
}
