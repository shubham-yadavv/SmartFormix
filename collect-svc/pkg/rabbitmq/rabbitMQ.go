package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue amqp.Queue
}

func NewRabbitMQService() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		"form",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{
		Conn:  conn,
		Ch:    ch,
		Queue: q,
	}, nil
}

func (r *RabbitMQ) PublishMessage(exchange string, key string, body []byte) error {
	err := r.Ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func PublishMessage(body []byte) error {
	rabbitMQ, err := NewRabbitMQService()
	if err != nil {
		return err
	}
	defer rabbitMQ.Conn.Close()
	defer rabbitMQ.Ch.Close()
	err = rabbitMQ.PublishMessage("", "form_responses", body)
	if err != nil {
		return err
	}
	return nil
}

// func (r *RabbitMQ) Consume(queue, consumer string, handlerFunc func([]byte) error) error {
// 	msgs, err := r.Ch.Consume(
// 		queue,
// 		consumer,
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	forever := make(chan bool)
// 	go func() {
// 		for d := range msgs {
// 			handlerFunc(d.Body)
// 		}
// 	}()
// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// 	return nil
// }

// func (r *RabbitMQ) Close() {
// 	r.Ch.Close()
// 	r.Conn.Close()
// }
