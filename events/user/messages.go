package events

import "time"

/*Construimos una estructura/interfaz para el mensaje que vamos a compartir entre servicios por
medio de la herramienta nats*/
/* We build a structure/interface for the message that we are going to share between services by
middle of the nats tool*/
type Message interface {
	Type() string
}

type CreatedUserMessage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (m CreatedUserMessage) Type() string {
	return "Hello_created_user"
}
