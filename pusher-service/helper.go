package main

import "time"

type CreatedUserMessage struct {
	Type      string    `json:"type"`
	ID        string    `json:"id"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func newCreatedUserMessage(id, address, email, name string, createdAt time.Time) *CreatedUserMessage {
	return &CreatedUserMessage{
		Type:      "created_user",
		ID:        id,
		Address:   address,
		Email:     email,
		Name:      name,
		CreatedAt: createdAt,
	}
}
