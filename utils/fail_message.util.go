package utils

import "log"

func FailMessage(err error, msg string) {
	if err == nil {
		return
	}
	log.Printf("%s: %s", msg, err.Error())
}
