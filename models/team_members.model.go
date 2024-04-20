package models

type TeamMembers struct {
	TeamId  string
	IdUsers []*string
}

type TeamsResponse struct {
	TeamMembers []*TeamMembers `json:"teams"`
}
