package models

type Template struct {
	IdTemplate  int    `json:"idTemplate"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}
