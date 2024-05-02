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
// ЕЛИСЕЙ: result service
// ЕЛИСЕЙ: проверить все запросы локально
// ЕЛИСЕЙ: потестить и посвязывать все с фронтом еще раз
// ЕЛИСЕЙ: проверка quest на время, если он прошел, то перестаем выдавать продумать как это сделать нормально
//
// ТИМУР: выгрузка в эксель опрос вопрос человекFrom человекTo ответ
//
// ЖДЕМ AUTH_SERVICE: гварды
// ЖДЕМ TEAM_SERVICE: полное создание опроса + полная выгрузка в эксель
//
// проверить документацию
// сделать прод версию
// провести мини уборку
// инверсия зависимостей?
// тесты?
//
// template и template_indicator изменить фк
//
// категории - работают, нельзя удалить если категория привязана к индикатору (так надо)
// индикаторы - работают, нельзя удалить если вопрос привязан к шаблону (так надо)
// шаблоны - работают, сейчас нельзя удалять из-за неправильной привязки фк (исправить)
// опросы - работают, удалить невозможно (поправить или не надо), проверить by-user (не работает)
// результаты - работает только создание (доделать)
//
// получение по айди опроса {команда, прогресс}
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
