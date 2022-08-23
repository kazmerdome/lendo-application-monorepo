POSTGRES_URI=postgresql://application_service@localhost:5432/application_service?sslmode=disable

migrate-up:
	migrate -path db/migration -database $(POSTGRES_URI) --verbose up

migrate-down:
	migrate -path db/migration -database $(POSTGRES_URI) --verbose down

build-application-job:
	go build -o build/application-job/application-job cmd/application-job/main.go

build-application-service:
	go build -o build/application-service/application-service cmd/application-service/main.go

.phony: migrate-up migrate-down build-application-job build-application-service

lint:
	golangci-lint run