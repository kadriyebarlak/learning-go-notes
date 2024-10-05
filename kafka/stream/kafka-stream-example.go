package main

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Consumer oluşturma
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// input-topic'ten mesajları al
	c.SubscribeTopics([]string{"input-topic"}, nil)

	// Producer'ı output-topic'e mesaj göndermek için oluştur
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	fmt.Println("consuming messages...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// Mesajı filtrele: Sadece "order" veya "purchase" içeren mesajları işleyelim
			messageValue := string(msg.Value)
			if strings.Contains(messageValue, "order") || strings.Contains(messageValue, "purchase") {
				fmt.Printf("İşlenen mesaj: %s\n", messageValue)

				// Filtrelenen mesajı output-topic'e gönder
				topic := "output-topic"
				p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          []byte(messageValue),
				}, nil)
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
