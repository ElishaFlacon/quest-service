package models

type LaunchQuest struct {
	IdLaunchQuest int  `json:"id"`
	IdTemplate    int  `json:"idTemplate"`
	IdTeam        int  `json:"idTeam"`
	Available     bool `json:"available"`
	StartAt       int  `json:"startAt"`
	EndAt         int  `json:"endAt"`
}
