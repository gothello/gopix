package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Open(uri string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}
