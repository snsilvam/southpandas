package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"southpandas.com/go/cqrs/database"
	events "southpandas.com/go/cqrs/events/user-client"
	repository "southpandas.com/go/cqrs/repository/user-client"
	search "southpandas.com/go/cqrs/search/user-client"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticsearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/users-clients", listUsersClientHandler).Methods(http.MethodGet)
	router.HandleFunc("/users-clients-search", searchHandler).Methods(http.MethodGet)
	return
}

func main() {
	//Procesa, las variables de entorno, definidas en el config
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//Conexion para el servicio de postgres
	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	repo, err := database.NewPostgresRepositoryUserClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetClientRepository(repo)
	//Conexion para el servicio de elastic search
	es, err := search.NewElastic(fmt.Sprintf("http://%s", cfg.ElasticsearchAddress))
	if err != nil {
		log.Fatal(err)
	}
	search.SetSearchRepository(es)
	defer search.Close()

	//Conexion con nats
	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		log.Fatal(err)
	}
	//Suscribimos nuestro query a un evento
	err = n.OnCreateUserClient(onCreatedUserClient)
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
