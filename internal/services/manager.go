package services

import (
	"chat-jobsity/internal/command"
	"chat-jobsity/internal/models"
	"fmt"
	"strings"
)

//go:generate mockgen -source=$GOFILE -package=mock_service -destination=../../test/mock/service/$GOFILE

type MessageManager interface {
	ManageMessage(m *models.MessageRequest) (string, error)
}

type MessageManagerImpl struct {
	cmdManager command.Manager
}

func NewMessageManager(cmdManager command.Manager) MessageManager {
	return &MessageManagerImpl{
		cmdManager: cmdManager,
	}
}

func (mv *MessageManagerImpl) ManageMessage(m *models.MessageRequest) (string, error) {
	isCommand := strings.HasPrefix(m.Text, "/")
	if isCommand {
		str := strings.Split(m.Text, "=")
		if len(str) < 2 {
			return fmt.Sprintf("ChatBot value is missing for %s", str[0]), nil
		}
		cmd := str[0]
		value := str[1]
		cmdRunner, err := mv.cmdManager.GetCommand(cmd)
		if err != nil {
			return "", err
		}

		msg, err := cmdRunner.Run(value)
		if err != nil {
			return "", err
		}
		m.Text = msg

		return "", nil
	}

	return m.ToBroadcast(), nil
}
