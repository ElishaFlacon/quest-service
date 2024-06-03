package utils

import (
	"fmt"
	"os"
)

func GetQuestUrl() string {
	return os.Getenv("QUEST_URL")
}

func GetQuestLink(id int) string {
	url := GetQuestUrl()
	link := fmt.Sprintf("%s/%d", url, id)
	return link
}
