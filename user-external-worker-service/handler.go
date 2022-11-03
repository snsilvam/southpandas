package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	events "southpandas.com/go/cqrs/events/user-external-worker"
	"southpandas.com/go/cqrs/models"
	repository "southpandas.com/go/cqrs/repository/user-external-worker"
)

// Modelo para escribir en base de datos
type createUserExternalWorkerRequest struct {
	ID                    string `json:"id"`
	ContractType          string `json:"contract_type"`
	WorkExperience        string `json:"work_experience"`
	WorkRemote            string `json:"work_remote"`
	Willingnesstravel     string `json:"willingnesstravel"`
	CurrentSalary         string `json:"current_salary"`
	ExpectedSalary        string `json:"expected_salary"`
	PossibilityOfRotation string `json:"possibility_of_rotation"`
	Profilelinkedln       string `json:"profile_linkedln"`
	Workarea              string `json:"workarea"`
	DescriptionWorkArea   string `json:"description_workarea"`
	User_id               string `json:"user_id"`
}

func createUserExternalWorkerHandler(w http.ResponseWriter, r *http.Request) {
	//Inicializamos el objeto createUserExternalWorkerRequest
	var req createUserExternalWorkerRequest
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
	userExternalWorker := models.UserExternalWorker{
		ID:                    id.String(),
		ContractType:          req.ContractType,
		WorkExperience:        req.WorkExperience,
		WorkRemote:            req.WorkRemote,
		Willingnesstravel:     req.Willingnesstravel,
		CurrentSalary:         req.CurrentSalary,
		ExpectedSalary:        req.ExpectedSalary,
		PossibilityOfRotation: req.PossibilityOfRotation,
		Profilelinkedln:       req.Profilelinkedln,
		Workarea:              req.Workarea,
		DescriptionWorkArea:   req.DescriptionWorkArea,
		User_id:               req.User_id,
		CreatedAt:             createdAt,
	}
	//Enviamos nuestro objeto al metodo Insert del paquete repository
	if err := repository.InsertUserExternalWorker(r.Context(), &userExternalWorker); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//publicar nuevo evento en nats
	if err := events.PublishCreatedUserExternalWorker(r.Context(), &userExternalWorker); err != nil {
		log.Printf("El sistema, no pudo publicar, el usuario tipo ExternalWorkere, creado: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userExternalWorker)
}
