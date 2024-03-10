package main

import (
	"github.com/ElishaFlacon/questionnaire-service/config"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/routes"
	"github.com/ElishaFlacon/questionnaire-service/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	config.Init(app).Start()
	pool := database.Init(utils.GetDatabaseString())
	routes.Init(app, pool).Start()

	app.Run(utils.GetAppUrl())
}
