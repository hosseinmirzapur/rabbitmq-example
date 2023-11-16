package broker

import (
	"fmt"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

type broker struct {
	connAddr string
}

func New(connAddr string) *broker {
	return &broker{
		connAddr: connAddr,
	}
}

func (b *broker) Connect() (*amqp.Connection, error) {
	log.Println("MQTT:", "trying to connect...")

	conn, err := amqp.Dial(b.connAddr)
	if err != nil {
		return nil, err
	}

	log.Println(strings.Repeat("*", 20), "MQTT Logs", strings.Repeat("*", 20))
	return conn, nil
}

func (b *broker) CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (b *broker) DeclareQueue(channel *amqp.Channel, name string) (*amqp.Queue, error) {
	queue, err := channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func (b *broker) Publish(ch *amqp.Channel, msg []byte, key, exchange string) error {
	err := ch.Publish(key, exchange, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})

	if err != nil {
		return err
	}

	fmt.Println("message published successfully")

	return nil
}

func (b *broker) GetQueueInfo(queue *amqp.Queue) {
	fmt.Println(strings.Repeat("#", 10), "Queue Info", strings.Repeat("#", 10))

	fmt.Printf("consumers count: %d\n", queue.Consumers)
	fmt.Printf("messages count: %d\n", queue.Messages)
	fmt.Printf("queue name: %s\n", queue.Name)
}

func (b *broker) Consume(channel *amqp.Channel, queue, consumer string) {
	msgs, err := channel.Consume(queue, consumer, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("waiting for messages...")

	<-forever
}
