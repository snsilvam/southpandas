package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"southpandas.com/go/cqrs/events"
	"southpandas.com/go/cqrs/models"
	"southpandas.com/go/cqrs/repository"
	"southpandas.com/go/cqrs/search"
)

func onCreatedUser(m events.CreatedUserMessage) {
	user := models.User{
		ID:        m.ID,
		Address:   m.Address,
		Email:     m.Email,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
	if err := search.IndexUser(context.Background(), user); err != nil {
		log.Printf("Se presento un error en la indexacion de la informacion: %v", err)
	}
}
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	users, err := repository.ListUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "Por favor ingresa un query", http.StatusBadRequest)
	}

	users, err := search.SearchUser(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
