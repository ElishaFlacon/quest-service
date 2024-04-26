package utils

import (
	"fmt"
	"strings"
	"time"
)

func GetStringDate(timestamp int) string {
	unixTime := time.Unix(int64(timestamp), 0)
	stringDate := unixTime.Format("01/02/2006")
	splitedDate := strings.Split(stringDate, "/")
	formatDate := fmt.Sprintf("%s.%s.%s", splitedDate[1], splitedDate[0], splitedDate[2])
	return formatDate
}
