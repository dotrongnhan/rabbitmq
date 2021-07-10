package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"dotrongnhan.com/common"
)

func main() {
	// creates rmq connection
	rmq := common.RabbitMQ{
		ConnectionString: "amqp://tfs:tfs-ocg@174.138.40.239:5672/",
	}
	rmq.CreateConnection()
	defer rmq.Close()

	ch := rmq.GetChannel()

	q, err := ch.QueueDeclare(
		"ocg", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	file, err := os.Open("political_thought_works_corpus.csv")
	fmt.Println(file)
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		body := scanner.Text()
		common.Publish(ch, q.Name, body)
	}
	file.Close()

	msgs := common.Consume(ch, q.Name)
	common.FailOnError(err, "Failed to register a consumer")

	// forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

}
