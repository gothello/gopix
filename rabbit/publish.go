package rabbit

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(conn *amqp.Connection, queue string, data interface{}) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return err
	}

	// confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	// if err := ch.Confirm(false); err != nil {
	// 	log.Fatalf("confirm.select destination: %s", err)
	// }

	// confirmed := <-confirms
	// fmt.Println(confirmed)

	if err != nil {
		log.Println("error publish confirm rabbitMQ")
	}

	return nil
}
