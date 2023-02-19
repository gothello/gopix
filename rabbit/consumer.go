package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitInputChan struct {
	Delivery amqp.Delivery
	Error    error
}

func Consumer(conn *amqp.Connection, name string, out chan RabbitInputChan) {

	ch, err := conn.Channel()
	if err != nil {
		out <- RabbitInputChan{
			Delivery: amqp.Delivery{},
			Error:    err,
		}
	}

	q, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		out <- RabbitInputChan{
			Delivery: amqp.Delivery{},
			Error:    err,
		}
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for m := range msgs {
		out <- RabbitInputChan{
			Delivery: m,
			Error:    nil,
		}
	}
}
