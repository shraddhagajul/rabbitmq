## RPC
When the Client starts up, it creates an anonymous exclusive callback queue.
For an RPC request, the Client sends a message with two properties:  reply_to, which is set to the callback queue and correlation_id, which is set to a unique value for every request.
The request is sent to an rpc_queue queue.
The RPC worker (aka: server) is waiting for requests on that queue. When a request appears, it does the job and sends a message with the result back to the Client, using the queue from the reply_to field.
The client waits for data on the callback queue. When a message appears, it checks the correlation_id property. If it matches the value from the request it returns the response to the application.

## Message properties

The AMQP 0-9-1 protocol predefines a set of 14 properties that go with a message. Most of the properties are rarely used, with the exception of the following:

persistent: Marks a message as persistent (with a value of true) or transient (false). You may remember this property from the second tutorial.
content_type: Used to describe the mime-type of the encoding. For example for the often used JSON encoding it is a good practice to set this property to: application/json.
reply_to: Commonly used to name a callback queue.
correlation_id: Useful to correlate RPC responses with requests.

## Callback queue
A client sends a request message and a server replies with a response message. In order to receive a response we need to send a 'callback' queue address with the request.