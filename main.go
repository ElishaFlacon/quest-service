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
// не забывай перепроверить все
// проверка quest на время, если он прошел, то перестаем выдавать продумать как это сделать нормально
// гварды
// quest-service на все пути в доке
// для post put добавить поля ввода в доке
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
