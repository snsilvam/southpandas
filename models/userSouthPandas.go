package models

import "time"

type UserSouthPandas struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	User_ID   string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
