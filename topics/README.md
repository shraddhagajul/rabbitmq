## Topics
Messages sent to a topic exchange can't have an arbitrary routing_key - it must be a list of words, delimited by dots. 
Two important special cases for binding keys:
(*) (star) can substitute for exactly one word.
(#) (hash) can substitute for zero or more words.

Topic exchange is powerful and can behave like other exchanges.

When a queue is bound with "#" (hash) binding key - it will receive all the messages, regardless of the routing key - like in fanout exchange.

When special characters "*" (star) and "#" (hash) aren't used in bindings, the topic exchange will behave just like a direct one.

Messages that don't match any bindings and will be lost.