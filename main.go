package main

import (
	"chat-jobsity/internal/handler"
	"chat-jobsity/internal/server"
)

func main() {
	messageHandler := handler.NewMessageHandler()
	api := server.NewAPI(messageHandler)
	api.Server.Logger.Fatal(api.Server.Start(":1323"))
}
