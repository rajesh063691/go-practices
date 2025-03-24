package publisher

import (
	"errors"
	"fmt"
	"net"

	"github.com/segmentio/kafka-go"
)

func CreateTopic(brokers []string, topic string) (err error) {
	// Create a topic
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		err = errors.New("could not connect to kafka")
		return
	}

	defer conn.Close()
	// Controller is the Leader
	controller, err := conn.Controller()
	if err != nil {
		err = errors.New("could not connect to controller")
		return
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, fmt.Sprint(controller.Port)))
	if err != nil {
		err = errors.New("could not connect to controller connection")
		return
	}

	defer controllerConn.Close()

	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = controllerConn.CreateTopics(topicConfig)
	if err != nil {
		err = errors.New("topic could not be connected")
		return
	}

	return nil
}
