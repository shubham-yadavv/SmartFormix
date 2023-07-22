package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/models"
)

type FormController struct {
}

func (fc FormController) CreateForm(c *gin.Context) {
	// var form models.Form
	// c.BindJSON(&form)
	// config.DB.Create(&form)

	// err := fc.rmq.Publish("", "hello", "text/plain", []byte("Form created"))
	// if err != nil {
	// 	log.Println(err)
	// }

	// c.JSON(200, gin.H{
	// 	"message": "Form created successfully",
	// 	"form":    form,
	// })

}

func (fc FormController) GetForm(c *gin.Context) {
	var form models.Form
	config.DB.First(&form, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Form fetched successfully",
		"form":    form,
	})

}

func (fc FormController) GetForms(c *gin.Context) {
	var forms []models.Form
	config.DB.Find(&forms)
	c.JSON(200, gin.H{
		"message": "Forms fetched successfully",
		"forms":   forms,
	})

}

func (fc FormController) UpdateForm(c *gin.Context) {
	var form models.Form
	config.DB.First(&form, c.Param("id"))
	c.BindJSON(&form)
	config.DB.Save(&form)
	c.JSON(200, gin.H{
		"message": "Form updated successfully",
		"form":    form,
	})

}

func (fc FormController) DeleteForm(c *gin.Context) {
	var form models.Form
	config.DB.First(&form, c.Param("id"))
	config.DB.Delete(&form)
	c.JSON(200, gin.H{
		"message": "Form deleted successfully",
	})

}
