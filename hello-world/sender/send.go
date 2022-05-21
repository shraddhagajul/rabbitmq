package main

import (
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

	body := "Hello Shraddha"

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body:  []byte(body)})
	util.FailOnError(err, "Failed to publish a message")
}

