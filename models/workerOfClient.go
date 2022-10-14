package models

import "time"

type WorkerOfClient struct {
	ID            string    `json:"id"`
	Description   string    `json:"description"`
	UserClient_ID string    `json:"userClient_id"`
	CreatedAt     time.Time `json:"created_at"`
}
