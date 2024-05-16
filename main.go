package main

import (
	"github.com/ElishaFlacon/quest-service/config"
	"github.com/ElishaFlacon/quest-service/database"
	_ "github.com/ElishaFlacon/quest-service/docs"
	"github.com/ElishaFlacon/quest-service/routes"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

// Вам дали задачу добавить функционал в этот микросервис? Жаль
// Буквально вся нужная информация о проекте находится в README.md

// TODO
//
// проверить документацию в конце
// проверить все запросы локально | 03.05 сделал
// потестить и посвязывать все с фронтом еще раз | 03.05 сделал частично
//
// опросы - by-user доделать
// результаты - доделать остатки
//
// получение презультатов о айди опроса {команда, прогресс, + участники}
// получение результатов членов команды по айди команды и опроса?
//
// нужны кастомные варианты ответов
//
// TODO

// @title			Quest Service API
// @description		Зашли почитать документацию? Жаль
// @host			localhost:5000
func main() {
	app := gin.Default()

	config.Init(app)
	database.Init(utils.GetDatabaseString())
	routes.Init(app)

	app.Run(utils.GetAppUrl())
}
