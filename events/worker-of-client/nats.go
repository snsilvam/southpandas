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
	conn                      *nats.Conn
	workerOfClientCreatedSub  *nats.Subscription
	workerOfClientCreatedChan chan CreatedWorkerOfClientMessage
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
	if n.workerOfClientCreatedSub != nil {
		n.workerOfClientCreatedSub.Unsubscribe()
	}
	close(n.workerOfClientCreatedChan)
}
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (n *NatsEventStore) PublishCreatedWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error {
	msg := CreatedWorkerOfClientMessage{
		ID:            workerOfClient.ID,
		Description:   workerOfClient.Description,
		UserClient_ID: workerOfClient.UserClient_ID,
		CreatedAt:     workerOfClient.CreatedAt,
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
func OnCreateWorkerOfClient(ctx context.Context, f func(CreatedWorkerOfClientMessage)) error {
	return eventStore.OnCreateWorkerOfClient(f)
}
func (n *NatsEventStore) OnCreateWorkerOfClient(f func(CreatedWorkerOfClientMessage)) (err error) {
	msg := CreatedWorkerOfClientMessage{}
	n.workerOfClientCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})
	return
}
func (n *NatsEventStore) SubscribeCreatedWorkerOfClient(ctx context.Context) (<-chan CreatedWorkerOfClientMessage, error) {
	m := CreatedWorkerOfClientMessage{}
	n.workerOfClientCreatedChan = make(chan CreatedWorkerOfClientMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.workerOfClientCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.workerOfClientCreatedChan <- m
			}
		}
	}()
	return (<-chan CreatedWorkerOfClientMessage)(n.workerOfClientCreatedChan), nil
}
