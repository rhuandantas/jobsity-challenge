package util

import (
	"chat-jobsity/internal/command"
	"chat-jobsity/internal/models"
	"errors"
	"strings"
)

type MessageValidator struct {
}

func NewMessageValidator() *MessageValidator {
	return &MessageValidator{}
}

func (mv *MessageValidator) ValidateMessage(m *models.MessageRequest) (string, error) {
	isCommand := strings.HasPrefix(m.Text, "/")
	if isCommand {
		str := strings.Split(m.Text, "=")
		if len(str) < 2 {
			return "", errors.New("command is invalid")
		}
		cmd := str[0]
		value := str[1]
		cmdRunner, err := command.GetCommand(cmd)
		if err != nil {
			return "", err
		}

		msg, err := cmdRunner.Run(value)
		if err != nil {
			return "", err
		}
		m.Text = msg
		return m.ToBroadcast(), nil
	}

	return m.ToBroadcast(), nil
}
