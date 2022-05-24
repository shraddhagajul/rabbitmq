package main

import (
	"github.com/streadway/amqp"
	"rabbitmq/util"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil)
	util.FailOnError(err, "Failed to declare an exchange")

	body := "info.error"
	err = ch.Publish("logs_topic", "info.error", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	util.FailOnError(err, "Failed to push a message to routing key info.error")

	body = "red.warning.*"
	err = ch.Publish("logs_topic", "red.warning.*", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	util.FailOnError(err, "Failed to push a message to routing key warning")

	body = "error"
	err = ch.Publish("logs_topic", "error", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	util.FailOnError(err, "Failed to push a message to routing key error")

	
	
}