package skill

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/sarama"
)

var (
	broker = os.Getenv("BROKER")
	topic  = os.Getenv("TOPIC")
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) SendMessageWithAction(action string, skill Skill) error {
	message := map[string]interface{}{
		"action": action,
		"data":   skill,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return err
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
