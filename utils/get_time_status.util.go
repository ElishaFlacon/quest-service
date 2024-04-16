package utils

import (
	"time"
)

func GetQuestTimeStatus(StartAt int, EndAt int) string {
	currentMinutes := time.Now().Minute()

	isReady := StartAt < currentMinutes
	isStart := StartAt >= currentMinutes && EndAt <= currentMinutes

	if isReady {
		return "Подготовлен"
	}

	if isStart {
		return "Запущен"
	}

	return "Завершен"

}
