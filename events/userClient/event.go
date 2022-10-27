package events

import (
	"context"

	"southpandas.com/go/cqrs/models"
)

/*Con el objetivo de hacer la abstracción, para implementar el principio de inversión de dependencias(SOLID),
creamos la siguiente interfaz*/
/*In order to make the abstraction, to implement the dependency inversion principle(SOLID),
we create the following interface */
type EventStore interface {
	Close()
	PublishCreatedUserClient(ctx context.Context, userClient *models.UserClient) error
	SubscribeCreatedUserClient(ctx context.Context) (<-chan CreatedUserClientMessage, error)
	OnCreateUserClient(f func(CreatedUserClientMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}
func PublishCreatedUserClient(ctx context.Context, userClient *models.UserClient) error {
	return eventStore.PublishCreatedUserClient(ctx, userClient)
}
func SubscribeCreatedUserClient(ctx context.Context) (<-chan CreatedUserClientMessage, error) {
	return eventStore.SubscribeCreatedUserClient(ctx)
}
