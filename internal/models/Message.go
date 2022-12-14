package models

import (
	"chat-jobsity/internal/command"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"name"`
	CreatedTS time.Time `json:"created_ts"`
}

func (m Message) ValidateMessage() (string, error) {
	isCommand := strings.HasPrefix(m.Text, "/")
	if isCommand {
		str := strings.Split(m.Text, "=")
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

		return msg, nil
	}

	return "", nil
}
