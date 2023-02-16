package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/collect-svc/pkg/routes"
)

func init() {
	// config.LoadEnvVariables()
	// config.ConnectDB()
	// config.SyncDatabase()

}

func main() {
	r := gin.Default()

	// rabbitmq.NewRabbitMQ()
	// Routes

	formRoutes := routes.FormRoutes{}
	formRoutes.FormRoutes(r)

	routes.TestRoutes(r)
	// controllers.ReceiveMessage()

	r.Run("localhost:3000")

}
