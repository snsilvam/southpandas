package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	events "southpandas.com/go/cqrs/events/worker-of-client"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/worker-of-client"
)

// Modelo para escribir en base de datos
type createWorkerOfClientRequest struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	UserClient_ID string `json:"userClient_id"`
}

func createWorkerOfClientHandler(w http.ResponseWriter, r *http.Request) {
	//Inicializamos el objeto createWorkerOfClientRequest
	var req createWorkerOfClientRequest
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
	WorkerOfClient := models.WorkerOfClient{
		ID:            id.String(),
		Description:   req.Description,
		UserClient_ID: req.UserClient_ID,
		CreatedAt:     createdAt,
	}
	//Enviamos nuestro objeto al metodo Insert del paquete repository
	if err := repository.InsertWorkerOfClient(r.Context(), &WorkerOfClient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//publicar nuevo evento en nats
	if err := events.PublishCreatedWorkerOfClient(r.Context(), &WorkerOfClient); err != nil {
		log.Printf("El sistema, no pudo publicar, el usuario tipo cliente, creado: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WorkerOfClient)
}
