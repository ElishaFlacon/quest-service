package models

type Result struct {
	IdResult      int    `json:"id_result"`
	IdIndicator   int    `json:"id_indicator"`
	IdLaunchQuest int    `json:"id_launch_quest"`
	IdFromUser    int    `json:"id_from_user"`
	IdToUser      int    `json:"id_to_user"`
	Value         string `json:"value"`
}
