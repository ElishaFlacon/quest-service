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

	config.Init(app)
	database.Init(utils.GetDatabaseString())
	routes.Init(app)

	app.Run(utils.GetAppUrl())
}
