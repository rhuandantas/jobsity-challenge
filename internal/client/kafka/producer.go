package kafka

import (
	"chat-jobsity/internal/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	cfgManager config.ConfigStore
}

type Message struct {
	Text    []byte
	Session string
}

func NewKafkaProducer(cfgManager config.ConfigStore) *Producer {
	return &Producer{
		cfgManager: cfgManager,
	}
}

func (kp *Producer) SendMessage(message string) {
	server := fmt.Sprintf("%s:%s", kp.cfgManager.Get("kafka.host"), kp.cfgManager.Get("kafka.port"))
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		fmt.Printf("producer -> %s", err.Error())
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("producer -> delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("producer -> delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "message_as_command"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	p.Flush(15 * 1000)
}
