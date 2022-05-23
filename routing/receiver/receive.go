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

	q, err := ch.QueueDeclare("info-error-msgs", false, false, true, false, nil)
	util.FailOnError(err, "Failed to declare a info-error-msgs queue")

	eq, err := ch.QueueDeclare("error-msgs", false, false, true, false, nil)
	util.FailOnError(err, "Failed to declare a error-msgs queue")

	err = ch.QueueBind(q.Name, "info", "logs_direct", false, nil)
	util.FailOnError(err, "Failed to bind queue "+q.Name)

	err = ch.QueueBind(q.Name, "error", "logs_direct", false, nil)
	util.FailOnError(err, "Failed to bind queue "+q.Name)

	err = ch.QueueBind(eq.Name, "error", "logs_direct", false, nil)
	util.FailOnError(err, "Failed to bind queue "+eq.Name)

	qmsgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range qmsgs {
			log.Printf(" q msgs [x] %s", d.Body)
		}
	}()

	eqMsgs, err := ch.Consume(eq.Name, "", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")
	go func() {
		for d := range eqMsgs {
			log.Printf(" eq msgs [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
