package models

type Indicator struct {
	IdIndicator int      `json:"idIndicator"`
	Name        string   `json:"name"`
	IdCategory  int      `json:"idСategory"`
	Description string   `json:"description"`
	Answers     []string `json:"answers"`
	FromRole    string   `json:"fromRole"`
	ToRole      string   `json:"toRole"`
	Available   bool     `json:"available"`
}

type IndicatorCreateRequest struct {
	IdCategory  int      `json:"idCategory"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Answers     []string `json:"answers"`
	FromRole    string   `json:"fromRole"`
	ToRole      string   `json:"toRole"`
}

type IndicatorWithCategoryName struct {
	IdIndicator  int      `json:"idIndicator"`
	IdCategory   int      `json:"idСategory"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Answers      []string `json:"answers"`
	FromRole     string   `json:"fromRole"`
	ToRole       string   `json:"toRole"`
	CategoryName string   `json:"categoryName"`
	Available    bool     `json:"available"`
}
