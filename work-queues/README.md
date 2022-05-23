## Work Queues
The main idea behind Work Queues (aka: Task Queues) is to avoid doing a resource-intensive task immediately and having to wait for it to complete.
We encapsulate a task as a message and send it to a queue. A worker process running in the background will pop the tasks and eventually execute the job. When you run many workers the tasks will be shared between them.

## Round robin dispatching 
Result by running 2 workers
Worker 1
2022/05/23 17:27:47  [*] Waiting for messages. To exit press CTRL+C
2022/05/23 17:27:59 Received a message: 1
2022/05/23 17:27:59 Received a message: 3
2022/05/23 17:27:59 Received a message: 5

Worker 2
2022/05/23 17:27:53  [*] Waiting for messages. To exit press CTRL+C
2022/05/23 17:27:59 Received a message: 2
2022/05/23 17:27:59 Received a message: 4

RabbitMQ will send each message to the next consumer, in sequence. On average every consumer will get the same number of messages. This way of distributing messages is called round-robin.

## Message ack
With our current code, once RabbitMQ delivers a message to the consumer it immediately marks it for deletion. In this case, if you kill a worker we will lose the message it was just processing. We'll also lose all the messages that were dispatched to this particular worker but were not yet handled.

What we want ?
If a worker dies, we'd like the task to be delivered to another worker.

RabbitMQ supports message acknowledgments. An ack(nowledgement) is sent back by the consumer to tell RabbitMQ that a particular message has been received, processed and that RabbitMQ is free to delete it.
<!-- d.Ack(false) -->

If a consumer dies (its channel is closed, connection is closed, or TCP connection is lost) without sending an ack, RabbitMQ will understand that a message wasn't processed fully and will re-queue it. If there are other consumers online at the same time, it will then quickly redeliver it to another consumer. That way you can be sure that no message is lost, even if the workers occasionally die.    

A timeout (30 minutes by default) is enforced on consumer delivery acknowledgement. This helps detect buggy (stuck) consumers that never acknowledge deliveries

## Message durability 
When RabbitMQ quits or crashes it will forget the queues and messages unless you tell it not to. Two things are required to make sure that messages aren't lost: we need to mark both the queue and messages as durable

At this point we're sure that the task_queue queue won't be lost even if RabbitMQ restarts. Now we need to mark our messages as persistent - by using the amqp.Persistent option amqp.Publishing takes.

DeliveryMode.  Transient means higher throughput but messages will not be
restored on broker restart.  The delivery mode of publishings is unrelated
to the durability of the queues they reside on.  Transient messages will
not be restored to durable queues, persistent messages will be restored to
durable queues and lost on non-durable queues during server restart.

## Fair displatch => prefetch count = 1
in a situation with two workers, when all odd messages are heavy and even messages are light, one worker will be constantly busy and the other one will do hardly any work. Well, RabbitMQ doesn't know anything about that and will still dispatch messages evenly.

This happens because RabbitMQ just dispatches a message when the message enters the queue. It doesn't look at the number of unacknowledged messages for a consumer. It just blindly dispatches every n-th message to the n-th consumer.

don't dispatch a new message to a worker until it has processed and acknowledged the previous one. Instead, it will dispatch it to the next worker that is not still busy.