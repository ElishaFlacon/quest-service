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
