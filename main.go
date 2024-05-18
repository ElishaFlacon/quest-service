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
// проверить все запросы локально
// потестить и посвязывать все с фронтом еще раз
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
