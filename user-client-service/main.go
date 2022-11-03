package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"southpandas.com/go/cqrs/database"
	events "southpandas.com/go/cqrs/events/userClient"
	repository "southpandas.com/go/cqrs/repository/user-client"
)

// Conexion con la base de datos
type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/users-clients", createUserClientHandler).Methods(http.MethodPost)
	return
}
func main() {
	//Inicializamos el objeto Config
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	//Luego de realizar una conexion con la base de datos exitosa, utilizamos el objeto addr para inicializar el objeto del repository
	repo, err := database.NewPostgresRepositoryUserClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	//Por medio del objeto repo inicializamos el objeto Repository
	repository.SetClientRepository(repo)
	//Iniciamos el objeto para el bus de nast
	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		log.Fatal(err)
	}
	events.SetEventStore(n)
	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
