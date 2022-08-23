package main

import (
	"fmt"
	"log"

	querier "github.com/kazmerdome/application-monorepo/db/sqlc"
	"github.com/kazmerdome/application-monorepo/pkg/actor/repository"
	"github.com/kazmerdome/application-monorepo/pkg/actor/rmq"
	"github.com/kazmerdome/application-monorepo/pkg/actor/server"
	"github.com/kazmerdome/application-monorepo/pkg/domain/application"
	applicationservice "github.com/kazmerdome/application-monorepo/pkg/service/application-service"
	"github.com/kazmerdome/application-monorepo/pkg/util/config"
)

func main() {
	var (
		dbHost = config.GetEnvString("POSTGRES_HOST", "localhost")
		dbPort = config.GetEnvString("POSTGRES_PORT", "5432")
		dbUsr  = config.GetEnvString("POSTGRES_USER", "application_service")
		dbPwd  = config.GetEnvString("POSTGRES_PASSOWRD", "")
		dbName = config.GetEnvString("POSTGRES_DB", "application_service")

		rabbitHost = config.GetEnvString("RABBIT_HOST", "localhost")
		rabbitPort = config.GetEnvString("RABBIT_PORT", "5672")
		rabbitUsr  = config.GetEnvString("RABBIT_USER", "guest")
		rabbitPwd  = config.GetEnvString("RABBIT_PASSOWRD", "guest")

		httpPort = config.GetEnvString("HTTP_PORT", "5050")
	)

	// Init rmq and postgres
	mq := rmq.NewRmq(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUsr, rabbitPwd, rabbitHost, rabbitPort))
	pg := repository.NewPostgresRespository(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsr, dbPwd, dbName))

	// Connect
	mq.Connect()
	defer mq.Disconnect()

	pg.Connect()
	defer pg.Disconnect()

	// Init domain
	q := querier.New(pg.GetDB())
	ar := application.NewRepository(q)
	as := application.NewService(ar)

	// Init service
	appService := applicationservice.NewService(mq, as)
	appHandler := applicationservice.NewHandler(appService)

	// Init and run server
	server := server.NewServer(httpPort, appHandler)
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
