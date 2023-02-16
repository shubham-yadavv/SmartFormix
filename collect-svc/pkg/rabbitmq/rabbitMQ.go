package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue amqp.Queue
}

func NewRabbitMQService(connString string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		"hello",
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

func (r *RabbitMQ) Publish(exchange, key, contentType string, body []byte) error {
	return r.Ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})
}

func (r *RabbitMQ) Consume(queue, consumer string, handlerFunc func([]byte) error) error {
	msgs, err := r.Ch.Consume(
		queue,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			handlerFunc(d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}

func (r *RabbitMQ) Close() {
	r.Ch.Close()
	r.Conn.Close()
}
