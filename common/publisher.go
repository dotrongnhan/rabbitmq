package common

import "github.com/streadway/amqp"

func Publish(ch *amqp.Channel, qName, body string) error {
	err := ch.Publish(
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
	return err
}
