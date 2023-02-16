package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/controllers"
)

type FormRoutes struct{}

func (fr FormRoutes) FormRoutes(incomingRoutes *gin.Engine) {
	formController := controllers.FormController{}

	incomingRoutes.POST("/forms", formController.CreateForm)
	incomingRoutes.GET("/forms/:id", formController.GetForm)
	incomingRoutes.GET("/forms", formController.GetForms)

}
