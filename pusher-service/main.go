package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"southpandas.com/go/cqrs/events"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	//variables de entorno
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	hub := NewHub()
	//Conectarnos a nats
	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		log.Fatal(err)
	}

	err = n.OnCreateUser(func(m events.CreatedUserMessage) {
		hub.Broadcast(newCreatedUserMessage(m.ID, m.Address, m.Email, m.Name, m.CreatedAt), nil)
	})
	if err != nil {
		log.Fatal(err)
	}
	events.SetEventStore(n)
	defer events.Close()

	go hub.Run()

	http.HandleFunc("/ws", hub.HandleWebSocket)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
