package application

import "github.com/google/uuid"

type Application struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Status    string    `json:"status"`
}

type CreateApplicationDTO struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

type ListApplicationDTO struct {
	Status string `json:"status"`
}

type UpdateApplicationDTO struct {
	Status *string `json:"status"`
}
