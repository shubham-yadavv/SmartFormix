package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/models"
)

func CreateQuestion(c *gin.Context) {
	var question models.Question
	c.BindJSON(&question)
	config.DB.Create(&question)
	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully", "question": question})
}

func GetQuestion(c *gin.Context) {
	var question models.Question
	config.DB.First(&question, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Question fetched successfully", "question": question})
}

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	config.DB.Find(&questions)
	c.JSON(http.StatusOK, gin.H{"message": "Questions fetched successfully", "questions": questions})
}
