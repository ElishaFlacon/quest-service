package main

import (
	"github.com/ElishaFlacon/quest-service/config"
	"github.com/ElishaFlacon/quest-service/database"
	_ "github.com/ElishaFlacon/quest-service/docs"
	"github.com/ElishaFlacon/quest-service/routes"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

// TODO
//
// проверить документацию 03.05 сделал, проверить еще раз 15.05
//
// ЕЛИСЕЙ: проверить все запросы локально | 03.05 сделал
// ЕЛИСЕЙ: потестить и посвязывать все с фронтом еще раз | 03.05 сделал частично
// ЕЛИСЕЙ: result service
//
// не ЖДЕМ AUTH_SERVICE: гварды или уже все равно
// не ЖДЕМ TEAM_SERVICE: полное создание опроса + полная выгрузка в эксель
//
// категории - работают, нельзя удалить если категория привязана к индикатору (так надо)
// индикаторы - работают, нельзя удалить если вопрос привязан к шаблону (так надо)
// шаблоны - работают, сейчас нельзя удалять из-за неправильной привязки фк (исправить)
// опросы - работают, удалить невозможно (поправить или не надо), проверить by-user (не работает)
// результаты - работает только создание (доделать)
//
// получение презультатов о айди опроса {команда, прогресс, + участники}
// получение результатов членов команды по айди команды и опроса?
//
// TODO

// @title			Quest Service API
// @description		Зашли почитать документацию? жаль.
// @host			localhost:5000
func main() {
	app := gin.Default()

	config.Init(app)
	database.Init(utils.GetDatabaseString())
	routes.Init(app)

	app.Run(utils.GetAppUrl())
}
