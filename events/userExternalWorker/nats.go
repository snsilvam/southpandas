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
	conn                          *nats.Conn
	userExternalWorkerCreatedSub  *nats.Subscription
	userExternalWorkerCreatedChan chan CreatedUserExternalWorkerMessage
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
	if n.userExternalWorkerCreatedSub != nil {
		n.userExternalWorkerCreatedSub.Unsubscribe()
	}
	close(n.userExternalWorkerCreatedChan)
}
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (n *NatsEventStore) PublishCreatedUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error {
	msg := CreatedUserExternalWorkerMessage{
		ID:                    userExternalWorker.ID,
		ContractType:          userExternalWorker.ContractType,
		WorkExperience:        userExternalWorker.WorkExperience,
		WorkRemote:            userExternalWorker.WorkRemote,
		Willingnesstravel:     userExternalWorker.Willingnesstravel,
		CurrentSalary:         userExternalWorker.CurrentSalary,
		ExpectedSalary:        userExternalWorker.ExpectedSalary,
		PossibilityOfRotation: userExternalWorker.PossibilityOfRotation,
		Profilelinkedln:       userExternalWorker.Profilelinkedln,
		Workarea:              userExternalWorker.Workarea,
		DescriptionWorkArea:   userExternalWorker.DescriptionWorkArea,
		User_id:               userExternalWorker.User_id,
		CreatedAt:             userExternalWorker.CreatedAt,
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
func OnCreateUserExternalWorker(ctx context.Context, f func(CreatedUserExternalWorkerMessage)) error {
	return eventStore.OnCreateUserExternalWorker(f)
}
func (n *NatsEventStore) OnCreateUserExternalWorker(f func(CreatedUserExternalWorkerMessage)) (err error) {
	msg := CreatedUserExternalWorkerMessage{}
	n.userExternalWorkerCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})
	return
}
func (n *NatsEventStore) SubscribeCreatedUserExternalWorker(ctx context.Context) (<-chan CreatedUserExternalWorkerMessage, error) {
	m := CreatedUserExternalWorkerMessage{}
	n.userExternalWorkerCreatedChan = make(chan CreatedUserExternalWorkerMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.userExternalWorkerCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.userExternalWorkerCreatedChan <- m
			}
		}
	}()
	return (<-chan CreatedUserExternalWorkerMessage)(n.userExternalWorkerCreatedChan), nil
}
