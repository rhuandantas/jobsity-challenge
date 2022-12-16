package handler

import (
	"chat-jobsity/internal/models"
	"chat-jobsity/internal/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type MessageHandler struct {
	validator *util.MessageValidator
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		validator: util.NewMessageValidator(),
	}
}

func (h *MessageHandler) HandleMessage(msg []byte) (string, error) {
	msgRequest := &models.MessageRequest{}
	err := json.Unmarshal(msg, msgRequest)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	message, err := h.validator.ValidateMessage(msgRequest)
	if err != nil {
		return message, err
	}

	return message, nil
}
