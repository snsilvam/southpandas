package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	events "southpandas.com/go/cqrs/events/user-client"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-client"
)

// Modelo para escribir en base de datos
type createUserClientRequest struct {
	ID      string `json:"id"`
	Premium string `json:"premium"`
	User_ID string `json:"user_id"`
}

func createUserClientHandler(w http.ResponseWriter, r *http.Request) {
	//Inicializamos el objeto createUserClientRequest
	var req createUserClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Autogeneramos un id y el atributo CreatedAt
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Asignamos los datos recibido al objeto que tenemos definido
	userClient := models.UserClient{
		ID:        id.String(),
		Premium:   req.Premium,
		User_ID:   req.User_ID,
		CreatedAt: createdAt,
	}
	//Enviamos nuestro objeto al metodo Insert del paquete repository
	if err := repository.InsertUserClient(r.Context(), &userClient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//publicar nuevo evento en nats
	if err := events.PublishCreatedUserClient(r.Context(), &userClient); err != nil {
		log.Printf("El sistema, no pudo publicar, el usuario tipo cliente, creado: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userClient)
}
