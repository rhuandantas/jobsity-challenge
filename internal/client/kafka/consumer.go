package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/olahol/melody"
	"time"
)

type Consumer struct {
	ws *melody.Melody
}

func NewKafkaConsumer(ws *melody.Melody) *Consumer {
	return &Consumer{
		ws: ws,
	}
}

func (kc *Consumer) ReadMessage() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "chat-consumer",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"message_as_command"}, nil)

	run := true

	for run {
		msg, err := c.ReadMessage(5 * time.Second)
		if err == nil {
			fmt.Printf("consumer -> reading on %s: %s\n", msg.TopicPartition, string(msg.Value))
			kc.ws.Broadcast(msg.Value)
		} else if !err.(kafka.Error).IsFatal() {
			fmt.Printf("consumer -> error: %v\n", err)
		}
	}

	c.Close()
}
