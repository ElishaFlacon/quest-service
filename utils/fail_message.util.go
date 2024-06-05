package utils

import "log"

func FailMessage(err error, msg string) {
	if err == nil {
		return
	}
	log.Panicf("%s: %s", msg, err.Error())
}
