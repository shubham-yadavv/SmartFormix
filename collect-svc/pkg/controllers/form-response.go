package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/models"
	"github.com/shubham/collect-svc/pkg/rabbitmq"
)

type FormResponse struct {
	FormID     uint   `json:"form_id"`
	FormName   string `json:"form_name"`
	FormDesc   string `json:"form_desc"`
	QuestionID uint   `json:"question_id"`
	Question   string `json:"question"`
	AnswerID   uint   `json:"answer_id"`
	Answer     string `json:"answer"`
}

func HandleFormResponse(c *gin.Context) {
	var formResponse FormResponse

	if err := c.ShouldBindJSON(&formResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON data"})
		return
	}

	responseJSON, err := json.Marshal(formResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert response to JSON"})
		return
	}

	rabbitmq.PublishMessage(responseJSON)

	saveFormResponse(&formResponse)

	c.JSON(http.StatusOK, gin.H{"message": "Response added to the queue and database successfully"})
}

func saveFormResponse(formResponse *FormResponse) {

	form := models.Form{
		FormID:   formResponse.FormID,
		FormName: formResponse.FormName,
		FormDesc: formResponse.FormDesc,
	}
	config.DB.Create(&form)

	question := models.Question{
		QuestionID:   formResponse.QuestionID,
		QuestionText: formResponse.Question,
		FormID:       formResponse.FormID,
	}
	config.DB.Create(&question)

	answer := models.Answer{
		AnswerID:   formResponse.AnswerID,
		AnswerText: formResponse.Answer,
		QuestionID: formResponse.QuestionID,
	}

	config.DB.Create(&answer)
}
