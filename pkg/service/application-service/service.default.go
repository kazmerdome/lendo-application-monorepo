package applicationservice

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"

	"github.com/kazmerdome/application-monorepo/pkg/actor/rmq"
	"github.com/kazmerdome/application-monorepo/pkg/domain/application"
)

type service struct {
	mq                      rmq.Rmq
	as                      application.Service
	applicationCreatedQueue string
}

func NewService(mq rmq.Rmq, as application.Service) *service {
	return &service{
		mq:                      mq,
		as:                      as,
		applicationCreatedQueue: "created.application",
	}
}

func (a *service) CreateApplication(w http.ResponseWriter, r *http.Request) {
	var dto application.CreateApplicationDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		a.jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)

		return
	}

	app, err := a.as.CreateApplication(dto)
	if err != nil {
		a.jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)

		return
	}

	log.Printf("new saved application: %v", app)

	// publish
	err = a.PublishToApplicationCreatedQueue(app)
	if err != nil {
		a.jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)

		return
	}

	a.jsonResponse(w, map[string]any{"application": app}, http.StatusCreated)
}

func (a *service) ListApplication(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	result, err := a.as.ListApplication(application.ListApplicationDTO{Status: status})
	if err != nil {
		a.jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)

		return
	}

	a.jsonResponse(w, map[string]any{"applications": result}, http.StatusOK)
}

func (a *service) PublishToApplicationCreatedQueue(app *application.Application) error {
	// create queue
	_, err := a.mq.GetChannel().QueueDeclare(a.applicationCreatedQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(map[string]any{
		"event_id": uuid.New(),
		"data":     app,
	})
	if err != nil {
		return err
	}

	err = a.mq.GetChannel().Publish(
		"",
		a.applicationCreatedQueue,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			ContentType:  "application/json",
			Body:         body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("published new application created event : %v", app)

	return nil
}

// jsonResponse encode an http json response with the provided response code.
func (a *service) jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error JSON response encoding. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Error JSON response writer. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
