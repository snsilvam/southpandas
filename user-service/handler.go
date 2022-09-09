package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"southpandas.com/go/cqrs/events"
	"southpandas.com/go/cqrs/models"
	"southpandas.com/go/cqrs/repository"
)

//escribir en base de datos
type createUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:        id.String(),
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		CreatedAt: createdAt,
	}

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
