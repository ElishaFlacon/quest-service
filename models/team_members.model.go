package models

type TeamMembers struct {
	teamId  string
	idUsers []*string
}

type TeamsResponse struct {
	TeamMembers []*TeamMembers `json:"teams"`
}
