package main

import (
	"log"
	"rabbitmq/hello-world/util"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("greetings", false, false, false, false, nil)
	util.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		} 
	}()
  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}