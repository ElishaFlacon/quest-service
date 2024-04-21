package models

type Result struct {
	IdResult    int    `json:"id"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  string `json:"idFromUser"`
	IdToUser    string `json:"idToUser"`
	Value       string `json:"value"`
}

type GetByUsersIdRequest struct {
	Users []struct{ Id string } `json:"users"`
}

type ResultCreateRequest struct {
	Results []struct {
		IdQuest     int    `json:"id_quest"`
		IdIndicator int    `json:"id_indicator"`
		IdFromUser  string `json:"id_from_user"`
		IdToUser    string `json:"id_to_user"`
		Value       string `json:"value"`
	} `json:"results"`
}
