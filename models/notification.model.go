package models

type Notification struct {
	Email string
	Link  string
}

type NotificationQueue struct {
	ConsumerEmail string
	Title         string
	Message       string
	Link          string
	ButtonName    string
}
