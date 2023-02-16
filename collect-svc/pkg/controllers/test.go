package controllers

// func PublishMessage(c *gin.Context) {
// 	rabbitmq := rabbitmq.NewRabbitMQ()
// 	err := rabbitmq.Publish("", "hello", "text/plain", []byte(" yooooooo!"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "Message published successfully",
// 	})

// }

// func ReceiveMessage() {
// 	rabbitmq := rabbitmq.NewRabbitMQ()

// 	err := rabbitmq.Consume("hello", "consumer", func(body []byte) error {
// 		log.Printf("Received a message: %s", body)
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
