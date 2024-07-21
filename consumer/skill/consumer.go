package skill

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type message struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Data   Skill  `json:"data"`
}

type Consumer struct {
	consumer      sarama.Consumer
	storage       storager
	topic         string
	actionHandler *ActionHandler
}

func NewConsumer(broker, topic string, db storager) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer:      consumer,
		storage:       db,
		topic:         topic,
		actionHandler: NewActionHandler(db),
	}, nil
}

func (c *Consumer) Consume() {
	partitionConsumer, err := c.consumer.ConsumePartition(c.topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln(err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message message
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			c.actionHandler.HandleAction(message)
			consumed++
			log.Printf("Consumed message: %s", msg.Value)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v", err)
		case <-signals:
			log.Println("Interrupt is detected")
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d messages", consumed)
}

func (c *Consumer) Close() {
	if err := c.consumer.Close(); err != nil {
		log.Printf("Failed to close consumer: %v", err)
	}
}
