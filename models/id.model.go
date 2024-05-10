package models

type Id[T comparable] struct {
	Id T `json:"id"`
}
