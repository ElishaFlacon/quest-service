package controllers

import (
	"github.com/ElishaFlacon/questionnaire-service/service"
	"github.com/gin-gonic/gin"
)

type TTemplate struct{}

var Template *TTemplate

func (*TTemplate) Create(context *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Indicators  []int  `json:"indicators"`
	}
	errBody := context.BindJSON(&body)

	if errBody != nil {
		context.JSON(400, gin.H{"error": errBody.Error()})
		return
	}

	data, count, err := service.Template.Create(body.Name, body.Description, body.Indicators)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	returningData := map[string]any{
		"id":              data[0].IdTemplate,
		"name":            data[0].Name,
		"description":     data[0].Description,
		"addedIndicators": count,
	}

	context.JSON(200, gin.H{"data": returningData})
}
