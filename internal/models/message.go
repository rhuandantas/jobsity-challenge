package models

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MessageRequest struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"name"`
	CreatedTS time.Time `json:"created_ts"`
}

func (mr MessageRequest) ToBroadcast() string {
	return fmt.Sprintf("%s %s", mr.UserName, mr.Text)
}

type MessageResponse struct {
	Message *string `json:"message"`
	Error   *string `json:"error"`
}
