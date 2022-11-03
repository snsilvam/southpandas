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
	PublishCreatedUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error
	SubscribeCreatedUserExternalWorker(ctx context.Context) (<-chan CreatedUserExternalWorkerMessage, error)
	OnCreateUserExternalWorker(f func(CreatedUserExternalWorkerMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}
func PublishCreatedUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error {
	return eventStore.PublishCreatedUserExternalWorker(ctx, userExternalWorker)
}
func SubscribeCreatedUserExternalWorker(ctx context.Context) (<-chan CreatedUserExternalWorkerMessage, error) {
	return eventStore.SubscribeCreatedUserExternalWorker(ctx)
}
