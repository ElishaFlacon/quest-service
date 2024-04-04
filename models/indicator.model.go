package models

type Indicator struct {
	IdIndicator int    `json:"idIndicator"`
	IdCategory  int    `json:"idCategory"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Visible     bool   `json:"visible"`
}

type IndicatorResponse struct {
	IdIndicator int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
