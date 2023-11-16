package main

import (
	"github.com/hosseinmirzapur/rmq-example/broker"
)

func RunProd(qName, msg string) {
	b := broker.New("amqp://guest:guest@localhost:5672")

	conn, err := b.Connect()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := b.CreateChannel(conn)
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := b.DeclareQueue(ch, qName)
	if err != nil {
		panic(err)
	}

	err = b.Publish(ch, []byte(msg), q.Name, "")
	if err != nil {
		panic(err)
	}

	b.GetQueueInfo(q)
}
