package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	events "southpandas.com/go/cqrs/events/user-client"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-client"
	search "southpandas.com/go/cqrs/search/user-client"
)

func onCreatedUserClient(m events.CreatedUserClientMessage) {
	userClient := models.UserClient{
		ID: m.ID,
	}
	if err := search.IndexUserClient(context.Background(), userClient); err != nil {
		log.Printf("Se presento un error en la indexacion de la informacion: %v", err)
	}
}
func listUsersClientHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	userClients, err := repository.ListUsersClients(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userClients)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "Por favor ingresa un query", http.StatusBadRequest)
	}

	userClients, err := search.SearchUserClient(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userClients)
}
