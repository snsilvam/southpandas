package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	events "southpandas.com/go/cqrs/events/worker-of-client"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/worker-of-client"
	search "southpandas.com/go/cqrs/search/worker-of-client"
)

func onCreatedWokerOfClient(m events.CreatedWorkerOfClientMessage) {
	workerOfClient := models.WorkerOfClient{
		ID:            m.ID,
		Description:   m.Description,
		UserClient_ID: m.UserClient_ID,
		CreatedAt:     m.CreatedAt,
	}
	if err := search.IndexWorkerOfClient(context.Background(), workerOfClient); err != nil {
		log.Printf("Se presento un error en la indexacion de la informacion: %v", err)
	}
}
func listworkerOfClientHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	workerOfClient, err := repository.ListWorkersOfClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workerOfClient)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "Por favor ingresa un query", http.StatusBadRequest)
	}

	workerOfClient, err := search.SearchWorkerOfClient(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workerOfClient)
}
