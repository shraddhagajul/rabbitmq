package main

import (
	"log"
	"rabbitmq/work-queues/util"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("worker", true, false, false, false, nil)
	util.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
)
	util.FailOnError(err, "Failed to set QoS")
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		} 
	}()
  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}