package kafka

import (
	"chat-jobsity/internal/config"
	"chat-jobsity/internal/logging"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/olahol/melody"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=mock_kafka -destination=../../../test/mock/client/kafka/$GOFILE

type Consumer interface {
	ReadMessage()
}

type consumer struct {
	ws  *melody.Melody
	cfg config.ConfigStore
	log logging.SimpleLogger
}

func NewKafkaConsumer(ws *melody.Melody, cfg config.ConfigStore, log logging.SimpleLogger) Consumer {
	return &consumer{
		ws:  ws,
		cfg: cfg,
		log: log,
	}
}

func (kc *consumer) ReadMessage() {
	hosts := fmt.Sprintf("%s:%s", kc.cfg.GetString("kafka.host"), kc.cfg.GetString("kafka.port"))
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": hosts,
		"group.id":          "chat-consumer",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{kc.cfg.GetString("kafka.topic")}, nil)

	run := true

	for run {
		msg, err := c.ReadMessage(5 * time.Second)
		if err == nil {
			kc.log.Info(fmt.Sprintf("consumer -> reading on %s: %s", msg.TopicPartition, msg.Value))
			kc.ws.Broadcast(msg.Value)
		} else if !err.(kafka.Error).IsFatal() {
			kc.log.Warn(fmt.Sprintf("consumer -> error: %v", err))
		}
	}

	c.Close()
}
