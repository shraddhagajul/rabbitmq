package util

import "log"

// FailOnError : helper function to check the return value for each amqp call
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
