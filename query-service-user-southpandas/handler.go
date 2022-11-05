package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	events "southpandas.com/go/cqrs/events/user-south-pandas"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-southpandas"
	search "southpandas.com/go/cqrs/search/user-south-pandas"
)

func onCreatedUserSouthPandas(m events.CreatedUserSouthPandasMessage) {
	userSouthPandas := models.UserSouthPandas{
		ID:        m.ID,
		Type:      m.Type_user,
		User_ID:   m.User_ID,
		CreatedAt: m.CreatedAt,
	}
	if err := search.IndexUserSouthPandas(context.Background(), userSouthPandas); err != nil {
		log.Printf("Se presento un error en la indexacion de la informacion: %v", err)
	}
}
func listUserSouthPandasHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	userSouthPandas, err := repository.ListUsersSouthPandas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userSouthPandas)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "Por favor ingresa un query", http.StatusBadRequest)
	}

	userSouthPandas, err := search.SearchUserSouthPandas(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userSouthPandas)
}
