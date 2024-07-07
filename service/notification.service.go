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

func convertToNotificationJSON(user *models.User, link string) []byte {
	title := "Опрос"
	message := "Привет! Тебе был назначен опрос, пройди пожалуйста :3"
	button := "Пройти опрос"

	notification := models.NotificationQueue{
		ConsumerEmail: user.Email,
		Title:         title,
		Message:       message,
		Link:          link,
		ButtonName:    button,
	}

	jsonMessage, err := json.Marshal(notification)
	if err != nil {
		log.Println("ERROR (failed marshal notification): ", err.Error())
		return nil
	}

	return jsonMessage
}

func (*TNotification) SendNotification(users []*models.User, link string) {
	amqpServerURL := utils.GetAMQPUrl()

	connectRabbitMQ, err := amqp091.Dial(amqpServerURL)
	if err != nil {
		log.Println("ERROR (failed connect to RabbitMQ): ", err.Error())
		return
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		log.Println("ERROR (failed open channel): ", err.Error())
		return
	}
	defer channelRabbitMQ.Close()

	q, err := channelRabbitMQ.QueueDeclare(
		utils.GetAMQPQueue(),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		log.Println("ERROR (failed declare queue): ", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, user := range users {
		notification := convertToNotificationJSON(user, link)

		err = channelRabbitMQ.PublishWithContext(ctx,
			utils.GetAMQPExchange(),
			q.Name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        notification,
			})
		if err != nil {
			log.Println("ERROR (failed publish message): ", err.Error())
			continue
		}

		log.Printf(" [x] Sent %s\n", notification)
	}
}
