package models

type Indicator struct {
	IdIndicator int    `json:"id"`
	Name        string `json:"name"`
	IdCategory  int    `json:"idСategory"`
	Description string `json:"description"`
	FromRole    string `json:"fromRole"`
	ToRole      string `json:"toRole"`
	Visible     bool   `json:"visible"`
}

type IndicatorCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	FromRole    string `json:"fromRole"`
	ToRole      string `json:"toRole"`
	Visible     bool   `json:"visible"`
	IdCategory  int    `json:"idCategory"`
}

type IndicatorWithCategoryName struct {
	IdIndicator  int    `json:"id"`
	IdCategory   int    `json:"idСategory"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FromRole     string `json:"fromRole"`
	ToRole       string `json:"toRole"`
	CategoryName string `json:"categoryName"`
	Visible      bool   `json:"visible"`
}
