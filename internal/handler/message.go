package handler

import (
	"chat-jobsity/internal/models"
	"chat-jobsity/internal/services"
	"encoding/json"
)

//go:generate mockgen -source=$GOFILE -package=mock_handler -destination=../../test/mock/handler/$GOFILE

type MessageHandler interface {
	HandleMessage(msg []byte) (string, error)
}

type messageHandler struct {
	manager services.MessageManager
}

func NewMessageHandler(manager services.MessageManager) MessageHandler {
	return &messageHandler{
		manager: manager,
	}
}

func (h *messageHandler) HandleMessage(msg []byte) (string, error) {
	msgRequest := &models.MessageRequest{}
	err := json.Unmarshal(msg, msgRequest)
	if err != nil {
		return "", err
	}

	return h.manager.ManageMessage(msgRequest)
}
