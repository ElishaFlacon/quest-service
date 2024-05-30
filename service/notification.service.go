package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/rabbitmq/amqp091-go"
)

type TNotification struct{}

var Notification *TNotification

func convertToNotificationJSON(receiver *models.Notification) []byte {
	notification := models.NotificationQueue{
		ConsumerEmail: receiver.Email,
		Title:         "Анонимный опрос",
		Message:       "Вам был назначен анонимный опрос оценки вашей команды. Перейдите по ссылке, чтобы пройти опрос",
		Link:          receiver.Link,
		ButtonName:    "Перейти по ссылке",
	}
	jsonMessage, err := json.Marshal(notification)
	utils.FailMessage(err, "Failed to convertToJson")
	return jsonMessage
}

func (*TNotification) SendNotification(receivers []*models.Notification) {
	amqpServerURL := utils.GetAMQPUrl()

	connectRabbitMQ, err := amqp091.Dial(amqpServerURL)
	utils.FailMessage(err, "Failed to connect to RabbitMQ")
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	utils.FailMessage(err, "Failed to open a channel")
	defer channelRabbitMQ.Close()

	q, err := channelRabbitMQ.QueueDeclare(
		utils.GetAMQPQueue(),
		true,
		false,
		false,
		true,
		nil,
	)
	utils.FailMessage(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, receiver := range receivers {
		notification := convertToNotificationJSON(receiver)
		err = channelRabbitMQ.PublishWithContext(ctx,
			utils.GetAMQPExchange(),
			q.Name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        notification,
			})
		utils.FailMessage(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", notification)
	}
}
