package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	events "southpandas.com/go/cqrs/events/user"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user"
)

// Modelo para escribir en base de datos
type createUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	//Inicializamos el objeto createUserRequest
	var req createUserRequest
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
	user := models.User{
		ID:        id.String(),
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		CreatedAt: createdAt,
	}
	//Enviamos nuestro objeto al metodo Insert del paquete repository
	if err := repository.InsertUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//publicar nuevo evento en nats
	if err := events.PublishCreatedUser(r.Context(), &user); err != nil {
		log.Printf("El sistema, no pudo publicar, el usuario creado: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
