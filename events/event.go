package events

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedUser(ctx context.Context, user *models.User) error
	SubscribeCreatedUser(ctx context.Context) (<-chan CreatedUserMessage, error)
	OnCreateUser(f func(CreatedUserMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}
func PublishCreatedUser(ctx context.Context, user *models.User) error {
	return eventStore.PublishCreatedUser(ctx, user)
}
func SubscribeCreatedUser(ctx context.Context) (<-chan CreatedUserMessage, error) {
	return eventStore.SubscribeCreatedUser(ctx)
}
