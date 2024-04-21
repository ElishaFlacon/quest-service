package models

type Template struct {
	IdTemplate  int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Available   bool    `json:"available"`
}

type TemplateWithCountIndicators struct {
	IdTemplate      int     `json:"id"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	Available       bool    `json:"available"`
	AddedIndicators int64   `json:"addedIndicators"`
}

type TemplateWithIndicators struct {
	IdTemplate  int         `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Available   bool        `json:"available"`
	Indicators  []Indicator `json:"indicators"`
}

type TemplateCreateRequest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Indicators  []struct{ Id int } `json:"indicators"`
}
