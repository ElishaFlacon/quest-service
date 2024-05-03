package models

type User struct {
	IdUser string `json:"idUser"`
	Status string `json:"status"`
}

type UsersInfo struct {
	IdResult         int    `json:"idResult"`
	FullNameToUser   string `json:"fullNameToUser"`
	FullNameFromUser string `json:"fullNameFromUser"`
}

type UsersInfoResponse struct {
	UsersInfo []*UsersInfo `json:"usersInfo"`
}

type UsersFromAndToByResultId struct {
	IdResult   int    `json:"idResult"`
	IdToUser   string `json:"idToUser"`
	IdFromUser string `json:"idFromUser"`
}
