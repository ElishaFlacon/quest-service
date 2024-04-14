package models

type TeamMembers struct {
	teamId  int
	idUsers []*int
}

type TeamsResponse struct {
	TeamMembers []*TeamMembers `json:"teams"`
}
