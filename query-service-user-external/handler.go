package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	events "southpandas.com/go/cqrs/events/user-external-worker"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-external-worker"
	search "southpandas.com/go/cqrs/search/user-external-worker"
)

func onCreatedUserExternalWorker(m events.CreatedUserExternalWorkerMessage) {
	userExternalWorker := models.UserExternalWorker{
		ID:                    m.ID,
		ContractType:          m.ContractType,
		WorkExperience:        m.WorkExperience,
		WorkRemote:            m.WorkRemote,
		Willingnesstravel:     m.Willingnesstravel,
		CurrentSalary:         m.CurrentSalary,
		ExpectedSalary:        m.ExpectedSalary,
		PossibilityOfRotation: m.PossibilityOfRotation,
		Profilelinkedln:       m.Profilelinkedln,
		Workarea:              m.Workarea,
		DescriptionWorkArea:   m.DescriptionWorkArea,
		User_id:               m.User_id,
		CreatedAt:             m.CreatedAt,
	}
	if err := search.IndexUserExternalWorker(context.Background(), userExternalWorker); err != nil {
		log.Printf("Se presento un error en la indexacion de la informacion: %v", err)
	}
}
func listUserExternalWorkerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	userExternalWorker, err := repository.ListUsersExternalWorkers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userExternalWorker)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "Por favor ingresa un query", http.StatusBadRequest)
	}

	userExternalWorker, err := search.SearchUserExternalWorker(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userExternalWorker)
}
