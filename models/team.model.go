package models

type User struct {
	IdUser string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserWithStatus struct {
	IdUser *string `json:"id"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Status bool    `json:"status"`
}

type UserWithStatusFull struct {
	IdUser  *string `json:"id"`
	IdTeam  string  `json:"idTeam"`
	IdQuest int     `json:"idQuest"`
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Status  bool    `json:"status"`
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

type TeamWithStatus struct {
	IdTeam  string            `json:"id"`
	Name    string            `json:"name"`
	Percent int               `json:"progress"`
	Users   []*UserWithStatus `json:"users"`
}

type TeamWithFullUser struct {
	IdTeam string      `json:"id"`
	Name   string      `json:"name"`
	Users  []*UserFull `json:"members"`
}
