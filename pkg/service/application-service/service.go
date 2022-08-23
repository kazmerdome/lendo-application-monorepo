package applicationservice

import (
	"net/http"

	"github.com/kazmerdome/application-monorepo/pkg/domain/application"
)

type Service interface {
	CreateApplication(w http.ResponseWriter, r *http.Request)
	ListApplication(w http.ResponseWriter, r *http.Request)
	PublishToApplicationCreatedQueue(app *application.Application) error
}
