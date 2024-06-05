package utils

import (
	"time"
)

func GetQuestTimeStatus(startAt int, endAt int) string {
	currentMinutes := time.Now().UnixMilli() / 1000

	startAt64 := int64(startAt)
	endAt64 := int64(endAt)

	isReady := startAt64 < currentMinutes
	isStart := startAt64 >= currentMinutes && endAt64 <= currentMinutes

	if isReady {
		return "Подготовлен"
	}

	if isStart {
		return "Запущен"
	}

	return "Завершен"

}
