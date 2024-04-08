package controllers

import (
	"strconv"

	"github.com/ElishaFlacon/quest-service/cruds"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/gin-gonic/gin"
)

type TTemplate struct{}

var Template *TTemplate

func (*TTemplate) Get(context *gin.Context) {
	params := context.Params
	strId, paramGood := params.Get("id")
	id, atoiErr := strconv.Atoi(strId)

	if !paramGood || atoiErr != nil {
		context.JSON(400, gin.H{"error": "Request params error!"})
		return
	}

	data, err := cruds.Template.Get(id)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data[0]})
}

func (*TTemplate) GetAll(context *gin.Context) {
	data, err := cruds.Template.GetAll()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TTemplate) Create(context *gin.Context) {
	var body struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Indicators  []struct{ Id int } `json:"indicators"`
	}
	errBody := context.BindJSON(&body)

	if errBody != nil {
		context.JSON(400, gin.H{"error": errBody.Error()})
		return
	}

	if len(body.Indicators) == 0 {
		context.JSON(400, gin.H{"error": "Indicators zero array"})
		return
	}

	indicators := []int{}
	for index := 0; index < len(body.Indicators); index++ {
		indicators = append(indicators, body.Indicators[index].Id)
	}

	data, count, err := service.Template.Create(body.Name, body.Description, indicators)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	returningData := models.TemplateWithCountIndicators{
		IdTemplate:      data[0].IdTemplate,
		Name:            data[0].Name,
		Description:     data[0].Description,
		Available:       data[0].Available,
		AddedIndicators: count,
	}

	context.JSON(200, gin.H{"data": returningData})
}
