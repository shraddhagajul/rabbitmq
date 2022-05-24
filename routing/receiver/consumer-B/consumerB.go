package main

import (
	"log"
	"rabbitmq/util"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp:guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)
	util.FailOnError(err, "Failed to declare an exchange")

	eQ, err := ch.QueueDeclare("error-msgs", false, false, true, false, nil)
	util.FailOnError(err, "Failed to declare a error-msgs queue")

	err = ch.QueueBind(eQ.Name, "error", "logs_direct", false, nil)
	util.FailOnError(err, "Failed to bind queue "+eQ.Name)

	forever := make(chan bool)

	eQMsgs, err := ch.Consume(eQ.Name, "", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")
	go func() {
		for d := range eQMsgs {
			log.Printf(" eQ msgs [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
