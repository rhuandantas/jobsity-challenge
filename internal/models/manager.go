package models

import (
	"chat-jobsity/internal/command"
	"fmt"
	"strings"
)

type MessageManager struct {
	cmdManager *command.Manager
}

func NewMessageManager(cmdManager *command.Manager) *MessageManager {
	return &MessageManager{
		cmdManager: cmdManager,
	}
}

func (mv *MessageManager) ManageMessage(m *MessageRequest) (string, error) {
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
