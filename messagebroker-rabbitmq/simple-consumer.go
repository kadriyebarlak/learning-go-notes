package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Consumer() {

	log.Printf("dialing %q", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		fmt.Printf("closing: %s", <-conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("declared Exchange, declaring Queue %q", "long-running-task-queue")
	queue, err := channel.QueueDeclare(
		"long-running-task-queue", // name of the queue
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // noWait
		nil,                       // arguments
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = channel.QueueBind(
		queue.Name,         // name of the queue
		"",                 // bindingKey
		"product-exchange", // sourceExchange
		false,              // noWait
		nil,                // arguments
	); err != nil {
		fmt.Println(err)
		return
	}

	deliveries, err := channel.Consume(
		queue.Name,         // name
		"product-consumer", // consumerTag,
		false,              // noAck
		false,              // exclusive
		false,              // noLocal
		false,              // noWait
		nil,                // arguments
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message := <-deliveries
		//async or long running task

		var newProduct Product
		json.Unmarshal(message.Body, &newProduct)
		if err != nil {
			log.Printf("Failed to unmarshal message: %s", err)
			continue
		}

		receivedMessages = append(receivedMessages, newProduct)
		log.Printf("Received message: %+v", newProduct)
		message.Ack(false)
	}

}
