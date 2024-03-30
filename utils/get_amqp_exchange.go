package utils

import (
	"os"
)

func GetAMQPExchange() string {
	return os.Getenv("AMQP_EXCHANGE")
}
