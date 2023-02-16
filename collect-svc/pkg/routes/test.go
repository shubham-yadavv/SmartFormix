package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/controllers"
)

func TestRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/test", controllers.PublishMessage)

}
