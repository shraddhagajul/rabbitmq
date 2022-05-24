package main

import (
	"rabbitmq/util"
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

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")

	timer := []int{1, 2, 3, 4, 5}

	for _, val := range timer {
		sleepTime := strconv.Itoa(val)
		err = ch.Publish("logs", "", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(sleepTime)})
		util.FailOnError(err, "Failed to publish a message")
	}

}
