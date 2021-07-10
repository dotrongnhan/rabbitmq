package common

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	ConnectionString string
	conn             *amqp.Connection
}

func (rmq *RabbitMQ) CreateConnection() {
	conn, err := amqp.Dial(rmq.ConnectionString)
	FailOnError(err, "Failed to connect to RabbitMQ")
	// set internal variable
	rmq.conn = conn
}

func (rmq *RabbitMQ) GetChannel() *amqp.Channel {
	ch, err := rmq.conn.Channel()
	FailOnError(err, "Cannot create channel")
	return ch
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (rmq *RabbitMQ) Close() {
	rmq.conn.Close()
}
