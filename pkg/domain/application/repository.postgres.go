package application

import (
	"context"

	"github.com/google/uuid"
	querier "github.com/kazmerdome/application-monorepo/db/sqlc"
)

type repository struct {
	queries querier.Querier
}

func NewRepository(queries querier.Querier) *repository {
	return &repository{
		queries: queries,
	}
}

func (r *repository) CreateOne(dto CreateApplicationDTO) (*Application, error) {
	arg := querier.CreateApplicationParams{
		ID:        uuid.New(),
		FirstName: dto.Firstname,
		LastName:  dto.Lastname,
	}

	a, err := r.queries.CreateApplication(context.TODO(), arg)
	if err != nil {
		return nil, err
	}

	app := Application{
		ID:        a.ID,
		Firstname: a.FirstName,
		Lastname:  a.LastName,
		Status:    a.Status,
	}

	return &app, nil
}

func (r *repository) ListApplication(status string) ([]Application, error) {
	ql, err := r.queries.ListApplications(context.TODO(), status)
	if err != nil {
		return nil, err
	}

	apps := []Application{}

	for i := range ql {
		qa := ql[i]

		apps = append(apps, Application{
			ID:        qa.ID,
			Firstname: qa.FirstName,
			Lastname:  qa.LastName,
			Status:    qa.Status,
		})
	}

	return apps, nil
}

func (r *repository) UpdateOne(id uuid.UUID, dto UpdateApplicationDTO) error {
	return nil
}
