package utils

import (
	"os"
)

func GetAMQPUrl() string {
	return os.Getenv("AMQP_URL_SERVER")
}
