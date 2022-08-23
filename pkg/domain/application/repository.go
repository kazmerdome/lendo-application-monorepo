package application

import "github.com/google/uuid"

type Repository interface {
	CreateOne(dto CreateApplicationDTO) (*Application, error)
	ListApplication(status string) ([]Application, error)
	UpdateOne(id uuid.UUID, dto UpdateApplicationDTO) error
}
