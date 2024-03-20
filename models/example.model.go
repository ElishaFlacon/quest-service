package models

type Example struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type ExampleBody struct {
	Value string `json:"value"`
}
