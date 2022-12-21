package main

import (
	"chat-jobsity/internal/client"
	"chat-jobsity/internal/client/kafka"
	"chat-jobsity/internal/command"
	"chat-jobsity/internal/config"
	"chat-jobsity/internal/handler"
	"chat-jobsity/internal/logging"
	"chat-jobsity/internal/server"
	"chat-jobsity/internal/services"
)

func main() {
	cfgManager, err := config.NewConfigManager()
	if err != nil {
		panic(err.Error())
	}

	logger := logging.NewSimpleLogger(cfgManager)
	kafkaProducer := kafka.NewKafkaProducer(cfgManager, logger)
	stooqClient := client.NewStooqClient(cfgManager, kafkaProducer)
	stockCommand := command.NewStooqCommand(stooqClient)
	cmdManager := command.NewCommandManager(stockCommand)
	messageManager := services.NewMessageManager(cmdManager)
	messageHandler := handler.NewMessageHandler(messageManager)
	api := server.NewAPI(messageHandler)

	kafkaConsumer := kafka.NewKafkaConsumer(api.Melody, cfgManager, logger)
	go kafkaConsumer.ReadMessage()

	api.Server.Logger.Fatal(api.Server.Start(":3001"))
}
