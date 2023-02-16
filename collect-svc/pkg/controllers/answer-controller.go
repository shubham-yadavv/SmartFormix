package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/models"
)

func CreateAnswer(c *gin.Context) {

	var answer models.Answer

	c.BindJSON(&answer)

	config.DB.Create(&answer)

	c.JSON(200, gin.H{
		"message": "Answer created successfully",
		"answer":  answer,
	})

}

func GetAnswer(c *gin.Context) {

	var answer models.Answer

	config.DB.First(&answer, c.Param("id"))

	c.JSON(200, gin.H{
		"message": "Answer fetched successfully",
		"answer":  answer,
	})

}

func GetAnswers(c *gin.Context) {

	var answers []models.Answer

	config.DB.Find(&answers)

	c.JSON(200, gin.H{
		"message": "Answers fetched successfully",
		"answers": answers,
	})

}
