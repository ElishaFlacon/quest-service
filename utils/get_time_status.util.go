package utils

import (
	"time"
)

func GetTimeStatus(StartAt int, EndAt int) string {
	currentTimeMinutes := time.Now().Minute()

	if StartAt < currentTimeMinutes {
		return "Подготовлен"
	}

	if StartAt >= currentTimeMinutes && EndAt <= currentTimeMinutes {
		return "Запущен"
	}

	return "Завершен"

}
