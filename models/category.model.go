package models

type Category struct {
	IdCategory int    `json:"idCategory"`
	Name       string `json:"name"`
}

type CategoryCreate struct {
	Name string `json:"name"`
}
