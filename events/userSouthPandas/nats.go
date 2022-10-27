package events

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/nats-io/nats.go"
	"southpandas.com/go/cqrs/models"
)

/*Realizamos la implementación de la abstracción(EventStore), en nuestro módulo de bajo nivel, para continuar con los
principios SOLID*/
/*We carry out the implementation of the abstraction (Event Store), in our low-level module, to continue with the
SOLID principless*/
type NatsEventStore struct {
	conn                       *nats.Conn
	userSouthPandasCreatedSub  *nats.Subscription
	userSouthPandasCreatedChan chan CreatedUserSouthPandasMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{conn: conn}, nil
}
func (n *NatsEventStore) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
	if n.userSouthPandasCreatedSub != nil {
		n.userSouthPandasCreatedSub.Unsubscribe()
	}
	close(n.userSouthPandasCreatedChan)
}
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (n *NatsEventStore) PublishCreatedUserSouthPandas(ctx context.Context, userSouthPandas *models.UserSouthPandas) error {
	msg := CreatedUserSouthPandasMessage{
		ID:        userSouthPandas.ID,
		Type_user: userSouthPandas.Type,
		User_ID:   userSouthPandas.User_ID,
		CreatedAt: userSouthPandas.CreatedAt,
	}
	data, err := n.encodeMessage(msg)
	if err != nil {
		return err
	}
	return n.conn.Publish(msg.Type(), data)
}
func (n *NatsEventStore) decodeMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
func OnCreateUserSouthPandas(ctx context.Context, f func(CreatedUserSouthPandasMessage)) error {
	return eventStore.OnCreateUserSouthPandas(f)
}
func (n *NatsEventStore) OnCreateUserSouthPandas(f func(CreatedUserSouthPandasMessage)) (err error) {
	msg := CreatedUserSouthPandasMessage{}
	n.userSouthPandasCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})
	return
}
func (n *NatsEventStore) SubscribeCreatedUserSouthPandas(ctx context.Context) (<-chan CreatedUserSouthPandasMessage, error) {
	m := CreatedUserSouthPandasMessage{}
	n.userSouthPandasCreatedChan = make(chan CreatedUserSouthPandasMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.userSouthPandasCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.userSouthPandasCreatedChan <- m
			}
		}
	}()
	return (<-chan CreatedUserSouthPandasMessage)(n.userSouthPandasCreatedChan), nil
}
