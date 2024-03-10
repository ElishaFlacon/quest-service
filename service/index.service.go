package service

import "github.com/ElishaFlacon/questionnaire-service/database"

type TInit struct {
	db *database.TInit
}

func Init(db *database.TInit) *TInit {
	return &TInit{db: db}
}
