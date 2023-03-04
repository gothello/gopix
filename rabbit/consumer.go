package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

//RabbitInputChan is one structs conteins type Dilevirey amq.Delivery and Error type error.

// used for type chan transport inputs or errors.
type RabbitInputChan struct {
	Delivery amqp.Delivery
	Error    error
}

//Consumer is a function that create on channel and queue to consume messages from AMQP channel.

//Consumer receive AMQP connection type *amqp.Connection.
//Consumer receive Queue name for consumer.
//Consumer receive chanInputChan to transport AMQP Delivery and error if exist.

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
