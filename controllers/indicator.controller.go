package controllers

import (
	"strconv"

	"github.com/ElishaFlacon/questionnaire-service/cruds"
	"github.com/ElishaFlacon/questionnaire-service/service"
	"github.com/gin-gonic/gin"
)

type TIndicator struct{}

var Indicator *TIndicator

func (*TIndicator) Get(context *gin.Context) {
	params := context.Params
	strId, paramGood := params.Get("id")
	id, atoiErr := strconv.Atoi(strId)

	if !paramGood || atoiErr != nil {
		context.JSON(400, gin.H{"error": "Request params error!"})
		return
	}

	data, err := cruds.Indicator.Get(id)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TIndicator) GetAll(context *gin.Context) {
	data, err := service.Indicator.GetAll()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TIndicator) GetByQuestId(context *gin.Context) {
	params := context.Params
	strId, paramGood := params.Get("id")
	id, atoiErr := strconv.Atoi(strId)

	if !paramGood || atoiErr != nil {
		context.JSON(400, gin.H{"error": "Request params error!"})
		return
	}

	data, err := service.Indicator.GetByQuestId(id)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TIndicator) Create(context *gin.Context) {
	var body struct {
		IdCategory  int    `json:"idCategory"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	errBody := context.BindJSON(&body)

	if errBody != nil {
		context.JSON(400, gin.H{"error": errBody.Error()})
		return
	}

	data, err := service.Indicator.Create(body.Name, body.Description, body.IdCategory)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}
