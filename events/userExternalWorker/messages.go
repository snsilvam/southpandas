package events

import "time"

/*Construimos una estructura/interfaz para el mensaje que vamos a compartir entre servicios por
medio de la herramienta nats*/
/* We build a structure/interface for the message that we are going to share between services by
middle of the nats tool*/
type Message interface {
	Type() string
}

type CreatedUserExternalWorkerMessage struct {
	ID                    string    `json:"id"`
	ContractType          string    `json:"contract_type"`
	WorkExperience        string    `json:"work_experience"`
	WorkRemote            string    `json:"work_remote"`
	Willingnesstravel     string    `json:"willingnesstravel"`
	CurrentSalary         string    `json:"current_salary"`
	ExpectedSalary        string    `json:"expected_salary"`
	PossibilityOfRotation string    `json:"possibility_of_rotation"`
	Profilelinkedln       string    `json:"profile_linkedln"`
	Workarea              string    `json:"workarea"`
	DescriptionWorkArea   string    `json:"description_workarea"`
	User_id               string    `json:"user_id"`
	CreatedAt             time.Time `json:"created_at"`
}

func (m CreatedUserExternalWorkerMessage) Type() string {
	return "Hello_created_userExternalWorker"
}
