package applicationjob

import "github.com/kazmerdome/application-monorepo/pkg/domain/application"

type Service interface {
	ConsumeFromApplicationCreatedQueue(done chan bool) error
	ProcessNewApplication(app application.Application)
}
