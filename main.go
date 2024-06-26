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
// Буквально вся нужная информация о микросервисе находится в README.md

// Дополнительная информация о микросервисе
// - По любым вопросам сюда "elishaflacon@gmail.com"
// - Проект простой, архитектура слоистая, но мы не используем
//   сложные абстракции, которые просто не нужны для такого микросервиса
// - Мы не используем ормки, по причине - "хотелось sql пописать :3"
// - Тут есть немного проблемных мест по оптимизации
// - Если вы обновили микросервис, перед релизом не забывайте:
//   -- Проверить документацию
//   -- Сбилдить локально и запустить
//   -- Сбилдить в докере и запустить
//   -- Прогнать тесты, какие? :3
// - Не забываем почитать README.md

// @title			Quest Service API
// @description		Зашли почитать документацию? Жаль
// @host			localhost:5000
func main() {
	config.InitDotenv()
	gin.SetMode(utils.GetAppMode())

	app := gin.Default()
	app.Use(gin.Recovery())

	config.Init(app)
	database.Init(utils.GetDatabaseString())
	routes.Init(app)

	println("APP WORK")

	app.Run(utils.GetAppUrl())
}
