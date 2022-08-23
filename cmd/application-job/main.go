package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kazmerdome/application-monorepo/pkg/actor/partnersapi"
	"github.com/kazmerdome/application-monorepo/pkg/actor/rmq"
	applicationjob "github.com/kazmerdome/application-monorepo/pkg/service/application-job"
	"github.com/kazmerdome/application-monorepo/pkg/util/config"
)

func main() {
	var (
		rabbitHost = config.GetEnvString("RABBIT_HOST", "localhost")
		rabbitPort = config.GetEnvString("RABBIT_PORT", "5672")
		rabbitUsr  = config.GetEnvString("RABBIT_USER", "guest")
		rabbitPwd  = config.GetEnvString("RABBIT_PASSOWRD", "guest")
		partnerAPI = config.GetEnvString("PARTNER_ADDRESS", "http://localhost:8000")
	)

	// Init actors
	pAPI := partnersapi.NewPartnersapi(partnerAPI)
	mq := rmq.NewRmq(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUsr, rabbitPwd, rabbitHost, rabbitPort))

	// Connect to rmq
	mq.Connect()
	defer mq.Disconnect()

	// Init services
	ajs := applicationjob.NewService(mq, pAPI)

	// Consume
	done := make(chan bool)
	go ajs.ConsumeFromApplicationCreatedQueue(done)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Stop consumer
	done <- true
}
