package pkg

import (
	"fmt"
	"message-loop/pkg/rpcdemo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

/**
	Code used from:
	https://github.com/confluentinc/confluent-kafka-go
**/

// ConsumeMessage - Consume message off of the topic
func ConsumeMessage() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	defer c.Close()

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"loop-topic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		fmt.Println("Received " + msg.String())
		rpcdemo.MakegRPCRequest()
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
