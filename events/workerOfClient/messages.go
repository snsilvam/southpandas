package events

import "time"

/*Construimos una estructura/interfaz para el mensaje que vamos a compartir entre servicios por
medio de la herramienta nats*/
/* We build a structure/interface for the message that we are going to share between services by
middle of the nats tool*/
type Message interface {
	Type() string
}

type CreatedWorkerOfClientMessage struct {
	ID            string    `json:"id"`
	Description   string    `json:"description"`
	UserClient_ID string    `json:"userClient_id"`
	CreatedAt     time.Time `json:"created_at"`
}

func (m CreatedWorkerOfClientMessage) Type() string {
	return "Hello_created_workerOfClient"
}
