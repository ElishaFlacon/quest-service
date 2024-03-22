package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func convertToNotificationJSON(receiver models.Notification) []byte {
	notification := models.NotificationQueue{
		ConsumerEmail: receiver.Email,
		Title:         "Анонимный опрос",
		Message:       "Вам был назначен анонимный опрос оценки вашей команды. Перейдите по ссылке, чтобы пройти опрос",
		Link:          receiver.Link,
		ButtonName:    "Перейти по ссылке",
	}
	jsonMessage, err := json.Marshal(notification)
	failOnError(err, "Failed to convertToJson")
	return jsonMessage
}

func SendNotification(receivers []models.Notification) {

	amqpServerURL := utils.GetAMQPUrl()

	connectRabbitMQ, err := amqp091.Dial(amqpServerURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	failOnError(err, "Failed to open a channel")
	defer channelRabbitMQ.Close()

	q, err := channelRabbitMQ.QueueDeclare(
		utils.GetAMQPQueue(), // queue name
		true,
		false,
		false,
		true,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := range receivers {
		notification := convertToNotificationJSON(receivers[i])
		err = channelRabbitMQ.PublishWithContext(ctx,
			utils.GetAMQPExchange(), // exchange
			q.Name,                  // routing key
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        notification,
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", notification)
	}
}
