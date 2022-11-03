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
	conn                  *nats.Conn
	userClientCreatedSub  *nats.Subscription
	userClientCreatedChan chan CreatedUserClientMessage
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
	if n.userClientCreatedSub != nil {
		n.userClientCreatedSub.Unsubscribe()
	}
	close(n.userClientCreatedChan)
}
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (n *NatsEventStore) PublishCreatedUserClient(ctx context.Context, userClient *models.UserClient) error {
	msg := CreatedUserClientMessage{
		ID:        userClient.ID,
		Premium:   userClient.Premium,
		User_ID:   userClient.User_ID,
		CreatedAt: userClient.CreatedAt,
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
func OnCreateUserClient(ctx context.Context, f func(CreatedUserClientMessage)) error {
	return eventStore.OnCreateUserClient(f)
}
func (n *NatsEventStore) OnCreateUserClient(f func(CreatedUserClientMessage)) (err error) {
	msg := CreatedUserClientMessage{}
	n.userClientCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})
	return
}
func (n *NatsEventStore) SubscribeCreatedUserClient(ctx context.Context) (<-chan CreatedUserClientMessage, error) {
	m := CreatedUserClientMessage{}
	n.userClientCreatedChan = make(chan CreatedUserClientMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.userClientCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.userClientCreatedChan <- m
			}
		}
	}()
	return (<-chan CreatedUserClientMessage)(n.userClientCreatedChan), nil
}
