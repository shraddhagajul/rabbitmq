package main

import (
	"rabbitmq/work-queues/util"
	"strconv"

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

	timer := []int{1, 2, 3, 4, 5}

	for _, val := range timer {
	sleepTime := strconv.Itoa(val)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body:  []byte(sleepTime), DeliveryMode: amqp.Persistent})
	util.FailOnError(err, "Failed to publish a message")
	}

}