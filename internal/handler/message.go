package handler

import (
	"chat-jobsity/internal/models"
	"encoding/json"
)

type MessageHandler struct {
	manager *models.MessageManager
}

func NewMessageHandler(manager *models.MessageManager) *MessageHandler {
	return &MessageHandler{
		manager: manager,
	}
}

func (h *MessageHandler) HandleMessage(msg []byte) (string, error) {
	msgRequest := &models.MessageRequest{}
	err := json.Unmarshal(msg, msgRequest)
	if err != nil {
		return "", err
	}

	return h.manager.ManageMessage(msgRequest)
}
