package common

import "github.com/streadway/amqp"

func Consume(ch *amqp.Channel, qName string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		qName, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	FailOnError(err, "Cannot consume from queue")
	return msgs
}
