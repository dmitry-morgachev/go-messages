package kafka

import (
	"log"
	"message-service/internal/service"
	"message-service/pkg/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(cfg *config.Config) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBroker,
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	err = c.SubscribeTopics([]string{cfg.KafkaTopic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s", string(msg.Value))
			processMessage(msg)
		} else {
			log.Printf("Consumer error: %v (%v)", err, msg)
		}
	}
}

func processMessage(msg *kafka.Message) {
	err := service.MarkMessageAsProcessed(string(msg.Value))
	if err != nil {
		log.Printf("Failed to mark message as processed: %s", err)
	} else {
		log.Printf("Message processed successfully")
	}
}
