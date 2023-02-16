package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/models"
)

func CreateResponse(c *gin.Context) {

	var response models.Response

	c.BindJSON(&response)

	config.DB.Create(&response)

	c.JSON(200, gin.H{
		"message":  "Response created successfully",
		"response": response,
	})

}

func GetResponse(c *gin.Context) {

	var response models.Response

	config.DB.First(&response, c.Param("id"))

	c.JSON(200, gin.H{
		"message":  "Response fetched successfully",
		"response": response,
	})

}

func GetResponses(c *gin.Context) {

	var responses []models.Response

	config.DB.Find(&responses)

	c.JSON(200, gin.H{
		"message":   "Responses fetched successfully",
		"responses": responses,
	})

}
