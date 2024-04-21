package models

type Category struct {
	IdCategory int    `json:"id"`
	Name       string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `json:"name"`
}
