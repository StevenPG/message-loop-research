package pkg

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

/**
	Code used from:
	https://github.com/confluentinc/confluent-kafka-go
**/

// ProduceMessage - Provide message as string to be sent to Kafka
func ProduceMessage(message string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	for t := range time.Tick(5 * time.Second) {
		go sendMessage(p, message, t)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}


func sendMessage(p *kafka.Producer, message string, t time.Time) {
	// Produce messages to topic (asynchronously)
	topic := "loop-topic"

	msg := message + t.String()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)
}