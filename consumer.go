package main

import "github.com/hosseinmirzapur/rmq-example/broker"

func RunCons(qName string) {
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

	b.Consume(ch, qName, "")
}
