package application

import "github.com/google/uuid"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (r *service) CreateApplication(dto CreateApplicationDTO) (*Application, error) {
	return r.repo.CreateOne(dto)
}

func (r *service) ListApplication(dto ListApplicationDTO) ([]Application, error) {
	return r.repo.ListApplication(dto.Status)
}

func (r *service) UpdateApplication(id uuid.UUID, dto UpdateApplicationDTO) error {
	return r.repo.UpdateOne(id, dto)
}
