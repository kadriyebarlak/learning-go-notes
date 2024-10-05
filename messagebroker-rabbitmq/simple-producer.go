package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ() *RabbitMQ {
	log.Printf("dialing %q", "amqp://guest:guest@localhost:5672/")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(fmt.Errorf("dial: %s", err))
		return nil
	}
	return &RabbitMQ{
		conn: connection,
	}
}

func (rabbit RabbitMQ) Publish(body []byte) {

	log.Printf("got Connection, getting Channel")
	channel, err := rabbit.conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := channel.ExchangeDeclare(
		"product-exchange", // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // noWait
		nil,                // arguments
	); err != nil {
		fmt.Println(err)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		"product-exchange", // publish to an exchange
		"",                 // routing to 0 or more queues
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		fmt.Println(err)
	}

}
