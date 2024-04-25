package utils

import "time"

func GetStringDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	stringDate := t.Format(time.UnixDate)
	return stringDate
}
