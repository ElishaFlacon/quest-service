package models

type Template struct {
	IdTemplate  int     `json:"idQuest"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Available   bool    `json:"available"`
}

type TemplateWithCountIndicators struct {
	IdTemplate      int     `json:"idQuest"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	Available       bool    `json:"available"`
	AddedIndicators int64   `json:"addedIndicators"`
}
type TemplateWithIndicators struct {
	IdTemplate  int         `json:"idQuest"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Available   bool        `json:"available"`
	Indicators  []Indicator `json:"indicators"`
}
