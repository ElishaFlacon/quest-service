package models

type Indicator struct {
	IdIndicator int    `json:"idIndicator"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Role        string `json:"role"`
	Visible     bool   `json:"visible"`
}
