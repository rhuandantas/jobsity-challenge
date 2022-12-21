package kafka

import (
	"chat-jobsity/internal/config"
	"chat-jobsity/internal/logging"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//go:generate mockgen -source=$GOFILE -package=mock_kafka -destination=../../../test/mock/client/kafka/$GOFILE

type Producer interface {
	SendMessage(message string)
}

type producer struct {
	cfgManager config.ConfigStore
	log        logging.SimpleLogger
}

type Message struct {
	Text    []byte
	Session string
}

func NewKafkaProducer(cfgManager config.ConfigStore, log logging.SimpleLogger) Producer {
	return &producer{
		cfgManager: cfgManager,
		log:        log,
	}
}

func (kp *producer) SendMessage(message string) {
	server := fmt.Sprintf("%s:%s", kp.cfgManager.Get("kafka.host"), kp.cfgManager.Get("kafka.port"))
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		kp.log.Info("producer -> %s", err.Error())
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					kp.log.Warn("producer -> delivery failed: %v\n", ev.TopicPartition)
				} else {
					kp.log.Info("producer -> delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := kp.cfgManager.GetString("kafka.topic")
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	p.Flush(15 * 1000)
}
