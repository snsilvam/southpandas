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
	conn            *nats.Conn
	userCreatedSub  *nats.Subscription
	userCreatedChan chan CreatedUserMessage
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
	if n.userCreatedSub != nil {
		n.userCreatedSub.Unsubscribe()
	}
	close(n.userCreatedChan)
}
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (n *NatsEventStore) PublishCreatedUser(ctx context.Context, user *models.User) error {
	msg := CreatedUserMessage{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
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
func OnCreateUser(ctx context.Context, f func(CreatedUserMessage)) error {
	return eventStore.OnCreateUser(f)
}
func (n *NatsEventStore) OnCreateUser(f func(CreatedUserMessage)) (err error) {
	msg := CreatedUserMessage{}
	n.userCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})
	return
}
func (n *NatsEventStore) SubscribeCreatedUser(ctx context.Context) (<-chan CreatedUserMessage, error) {
	m := CreatedUserMessage{}
	n.userCreatedChan = make(chan CreatedUserMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.userCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.userCreatedChan <- m
			}
		}
	}()
	return (<-chan CreatedUserMessage)(n.userCreatedChan), nil
}
