package models

type Template struct {
	IdTemplate  int     `json:"idQuestTemplate"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Available   bool    `json:"available"`
}

type TemplateWithCountIndicators struct {
	IdTemplate      int     `json:"idQuestTemplate"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	Available       bool    `json:"available"`
	AddedIndicators int64   `json:"addedIndicators"`
}

type TemplateWithIndicators struct {
	IdTemplate  int                      `json:"idQuestTemplate"`
	Name        string                   `json:"name"`
	Description *string                  `json:"description"`
	Available   bool                     `json:"available"`
	Indicators  []*IndicatorWithCategory `json:"indicators"`
}

type TemplateCreate struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Indicators  []Id[int] `json:"indicators"`
}
