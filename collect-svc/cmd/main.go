package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/shubham/collect-svc/pkg/config"
	"github.com/shubham/collect-svc/pkg/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDatabase()
}

func main() {

	router := gin.Default()

	routes.TestRoutes(router)

	log.Println("HTTP server listening on port 3000...")
	log.Fatal(router.Run(":3000"))
}
