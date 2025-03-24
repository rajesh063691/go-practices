package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	brokers := []string{"localhost:29092"}
	topic := "TEST_TOPIC"
	groupID := "consumerGroup"

	// Create a context
	ctx := context.Background()

	// initialise the new reader with groupID and brokers
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()
	fmt.Println("Starting consuming...")

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error while reading message :: ", err)
			break
		}
		fmt.Printf("Message on %s: %s\n", msg.Topic, string(msg.Value))

		// simulate a long running process
		time.Sleep(2 * time.Second)
	}

}
