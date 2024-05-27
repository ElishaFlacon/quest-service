package models

type Result struct {
	IdResult    int    `json:"idResult"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  string `json:"idFromUser"`
	IdToUser    string `json:"idToUser"`
	Value       string `json:"value"`
}

type ResultWithInfo struct {
	IdResult    int    `json:"idResult"`
	IdIndicator int    `json:"idIndicator"`
	IdQuest     int    `json:"idQuest"`
	IdFromUser  string `json:"idFromUser"`
	IdToUser    string `json:"idToUser"`
	Value       string `json:"value"`

	Teams struct {
		IdTeam  string       `json:"idTeam"`
		Percent int          `json:"percent"`
		Users   []Id[string] `json:"users"`
	} `json:"teams"`
}

type GetByUsersIdRequest struct {
	Users []Id[string] `json:"users"`
}

type ResultResponse struct {
	CreatedResults int64 `json:"createdResults"`
}

type ResultCreateRequest struct {
	Results []struct {
		IdQuest     int    `json:"idQuest"`
		IdIndicator int    `json:"idIndicator"`
		IdFromUser  string `json:"idFromUser"`
		IdToUser    string `json:"idToUser"`
		Value       string `json:"value"`
	} `json:"results"`
}
