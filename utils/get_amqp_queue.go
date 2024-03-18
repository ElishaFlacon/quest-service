package utils

import (
	"os"
)

func GetAMQPQueue() string {
	return os.Getenv("AMQP_QUEUE_NAME")
}
