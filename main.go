package main

import (
	"github.com/ElishaFlacon/quest-service/config"
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/routes"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

// fromRole toRole for indicators
// перепроверить все
// docs
// взаимодействие слоев
// переписать форы на for := range

func main() {
	app := gin.Default()

	config.Init(app)
	database.Init(utils.GetDatabaseString())
	routes.Init(app)

	app.Run(utils.GetAppUrl())
}
