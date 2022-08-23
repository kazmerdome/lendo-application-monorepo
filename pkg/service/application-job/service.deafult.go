package applicationjob

import (
	"encoding/json"
	"log"

	"github.com/kazmerdome/application-monorepo/pkg/actor/partnersapi"
	"github.com/kazmerdome/application-monorepo/pkg/actor/rmq"
	"github.com/kazmerdome/application-monorepo/pkg/domain/application"
)

type service struct {
	mq                      rmq.Rmq
	pAPI                    partnersapi.Partnersapi
	applicationCreatedQueue string
}

func NewService(mq rmq.Rmq, pAPI partnersapi.Partnersapi) *service {
	return &service{
		mq:                      mq,
		pAPI:                    pAPI,
		applicationCreatedQueue: "created.application",
	}
}

func (r *service) ConsumeFromApplicationCreatedQueue(done chan bool) error {
	chnl := r.mq.GetChannel()

	// create queue
	_, err := chnl.QueueDeclare(r.applicationCreatedQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// consume
	msgs, err := chnl.Consume(
		r.applicationCreatedQueue,
		"partner-job",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// listen
	go func() {
		for d := range msgs {
			var event struct {
				Data application.Application `json:"data"`
			}

			if err = json.Unmarshal(d.Body, &event); err != nil {
				log.Panic(err)
			}

			log.Printf("New event received : %v", event)
			r.ProcessNewApplication(event.Data)
		}
	}()
	<-done
	log.Printf("closing ApplicationCreatedQueue consumer")
	return nil
}

func (r *service) ProcessNewApplication(app application.Application) {
	b, err := json.Marshal(app)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("sending new application sent to partner: %v ", app)

	resp, err := r.pAPI.PostApplication(b)
	if err != nil {
		log.Panic(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&app); err != nil {
		log.Panic(err)
	}

	log.Printf("new application sent to partner: %v ", app)
}
