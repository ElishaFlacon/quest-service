package models

type Indicator struct {
	IdIndicator int      `json:"id"`
	Name        string   `json:"name"`
	IdCategory  int      `json:"idСategory"`
	Description string   `json:"description"`
	Answers     []string `json:"answers"`
	FromRole    string   `json:"role"`
	ToRole      string   `json:"type"`
	Available   bool     `json:"available"`
}

type IndicatorWithCategory struct {
	IdIndicator  int      `json:"id"`
	IdCategory   int      `json:"idСategory"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Answers      []string `json:"answers"`
	FromRole     string   `json:"role"`
	ToRole       string   `json:"type"`
	CategoryName string   `json:"categoryName"`
	Available    bool     `json:"available"`
}

type IndicatorCreate struct {
	IdCategory  int      `json:"idCategory"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Answers     []string `json:"answers"`
	FromRole    string   `json:"fromRole"`
	ToRole      string   `json:"type"`
}
