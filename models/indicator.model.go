package models

type Indicator struct {
	IdIndicator int    `json:"id_indicator"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Role        string `json:"role"`
	Visible     bool   `json:"visible"`
}
