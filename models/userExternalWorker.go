package models

import "time"

type UserExternalWorker struct {
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
