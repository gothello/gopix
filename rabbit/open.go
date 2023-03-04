package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)


//This function is used to open a connection to an AMQP server using the provided URI. It begins by dialing the server using the amqp.Dial() function, which returns a connection and an error if one occurs. If an error occurs, it is logged and the function returns nil for the channel and the error. 

// no error occurs, a channel is created from the connection using conn.Channel(). If an error occurs during this step, it is returned as part of the return value of the function. Otherwise, a pointer to the channel is returned along with nil for the error.



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
