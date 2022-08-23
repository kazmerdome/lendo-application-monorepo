package application

import "github.com/google/uuid"

type Service interface {
	CreateApplication(dto CreateApplicationDTO) (*Application, error)
	ListApplication(dto ListApplicationDTO) ([]Application, error)
	UpdateApplication(id uuid.UUID, dto UpdateApplicationDTO) error
}
