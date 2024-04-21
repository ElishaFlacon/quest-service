package models

type Indicator struct {
	IdIndicator int    `json:"id"`
	Name        string `json:"name"`
	IdCategory  int    `json:"idСategory"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Visible     bool   `json:"visible"`
}

type IndicatorCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Visible     bool   `json:"visible"`
	IdCategory  int    `json:"idCategory"`
}

type IndicatorWithCategoryName struct {
	IdIndicator  int    `json:"id"`
	IdCategory   int    `json:"idСategory"`
	CategoryName int    `json:"categoryName"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Role         string `json:"role"`
	Visible      bool   `json:"visible"`
}
