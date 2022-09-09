package models

import "time"

type UserClient struct {
	ID        string    `json:"id"`
	Premium   bool      `json:"premium"`
	User_ID   string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
