package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	events "southpandas.com/go/cqrs/events/user-south-pandas"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-southpandas"
)

// Modelo para escribir en base de datos
type createUserSouthPandasRequest struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	User_ID string `json:"user_id"`
}

func createUserSouthPandasHandler(w http.ResponseWriter, r *http.Request) {
	//Inicializamos el objeto createUserSouthPandasRequest
	var req createUserSouthPandasRequest
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
	userSouthPandas := models.UserSouthPandas{
		ID:        id.String(),
		Type:      req.Type,
		User_ID:   req.User_ID,
		CreatedAt: createdAt,
	}
	//Enviamos nuestro objeto al metodo Insert del paquete repository
	if err := repository.InsertUserSouthPandas(r.Context(), &userSouthPandas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//publicar nuevo evento en nats
	if err := events.PublishCreatedUserSouthPandas(r.Context(), &userSouthPandas); err != nil {
		log.Printf("El sistema, no pudo publicar, el usuario tipo SouthPandase, creado: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userSouthPandas)
}
