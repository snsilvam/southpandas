package events

import "time"

type Message interface {
	Type() string
}

type CreatedUserMessage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (m CreatedUserMessage) Type() string {
	return "Hello_created_user"
}
