package main

import (
	"log"
	"os"

	"github.com/narunart-atise/skill-api-kafka/consumer/database"
	"github.com/narunart-atise/skill-api-kafka/consumer/skill"
)

func main() {
	broker := os.Getenv("BROKER")
	topic := os.Getenv("TOPIC")

	db, closeDB := database.NewPostgres()
	defer closeDB()

	storage := skill.NewStorage(db)

	consumer, err := skill.NewConsumer(broker, topic, storage)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	consumer.Consume()
}
