package models

type Id[T any] struct {
	Id T `json:"id"`
}
