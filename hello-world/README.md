## RabbitMQ
RabbitMQ is a message broker: it accepts and forwards messages.

RabbitMQ, and messaging in general, uses some jargon.
1. Producing => sending. A program that sends messages is a producer 
2. A queue lives inside RabbitMQ. Although messages flow through RabbitMQ and your applications, they can only be stored inside a queue. A queue is only bound by the host's memory & disk limits, it's essentially a large message buffer. Many producers can send messages that go to one queue, and many consumers can try to receive data from one queue
3. Consuming => receiving. A consumer is a program that mostly waits to receive messages

