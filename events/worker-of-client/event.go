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
	PublishCreatedWorkerOfClient(ctx context.Context, WorkerOfClient *models.WorkerOfClient) error
	SubscribeCreatedWorkerOfClient(ctx context.Context) (<-chan CreatedWorkerOfClientMessage, error)
	OnCreateWorkerOfClient(f func(CreatedWorkerOfClientMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}
func PublishCreatedWorkerOfClient(ctx context.Context, WorkerOfClient *models.WorkerOfClient) error {
	return eventStore.PublishCreatedWorkerOfClient(ctx, WorkerOfClient)
}
func SubscribeCreatedWorkerOfClient(ctx context.Context) (<-chan CreatedWorkerOfClientMessage, error) {
	return eventStore.SubscribeCreatedWorkerOfClient(ctx)
}
