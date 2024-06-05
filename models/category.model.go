package models

type Category struct {
	IdCategory int    `json:"idCategory"`
	Name       string `json:"name"`
	Available  bool   `json:"available"`
}

type CategoryCreate struct {
	Name string `json:"name"`
}
