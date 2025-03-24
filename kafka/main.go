package main

import (
	"context"
	"fmt"
	"kafka/publisher"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "TEST_TOPIC"

	fmt.Println(brokers, topic)

	err := publisher.CreateTopic(brokers, topic)
	if err != nil {
		log.Fatalf(err.Error(), err)
		return
	}

	fmt.Println("Topic created successfully :: ", topic)

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		err = writer.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte(fmt.Sprintf("Hello World! %d", i)),
			})
		if err != nil {
			log.Fatalf(err.Error(), err)
			return
		}

		fmt.Println("Message written successfully :: ", i)
	}

	fmt.Println("All Message published successfully to topic :: ", topic)

}
