package models

type User struct {
	IdUser string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserFull struct {
	IdUser    string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Team struct {
	IdTeam string  `json:"id"`
	Name   string  `json:"name"`
	Users  []*User `json:"members"`
}

type TeamWithFullUser struct {
	IdTeam string      `json:"id"`
	Name   string      `json:"name"`
	Users  []*UserFull `json:"members"`
}
