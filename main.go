package main

import (
	"flag"
)

func main() {
	producer := flag.Bool("producer", true, "RabbitMQ Producer")
	queueName := flag.String("qname", "testing", "RabbitMQ working queue name")
	msg := flag.String("msg", "test message", "user message")

	flag.Parse()

	if *producer {
		RunProd(*queueName, *msg)
	} else {
		RunCons(*queueName)
	}
}
