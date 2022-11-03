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
	PublishCreatedUserSouthPandas(ctx context.Context, userSouthPandas *models.UserSouthPandas) error
	SubscribeCreatedUserSouthPandas(ctx context.Context) (<-chan CreatedUserSouthPandasMessage, error)
	OnCreateUserSouthPandas(f func(CreatedUserSouthPandasMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}
func PublishCreatedUserSouthPandas(ctx context.Context, userSouthPandas *models.UserSouthPandas) error {
	return eventStore.PublishCreatedUserSouthPandas(ctx, userSouthPandas)
}
func SubscribeCreatedUserSouthPandas(ctx context.Context) (<-chan CreatedUserSouthPandasMessage, error) {
	return eventStore.SubscribeCreatedUserSouthPandas(ctx)
}
