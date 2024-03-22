package models

type LaunchQuest struct {
	IdLaunchQuest int  `json:"id_launch_quest"`
	IdQuest       int  `json:"id_quest"`
	IdTeam        int  `json:"id_team"`
	Available     bool `json:"available"`
	StartAt       int  `json:"start_at"`
	EndAt         int  `json:"end_at"`
}
