package models

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	Messages  []Message `json:"messages"`
	CreatedTS time.Time `json:"created_ts"`
}
