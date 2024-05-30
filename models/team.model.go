package models

type User struct {
	IdUser     string `json:"id"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

type Team struct {
	IdTeam string  `json:"id"`
	Users  []*User `json:"members"`
}
