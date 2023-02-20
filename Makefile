all: test migrate generate run-local

generate:
	buf generate

migrate:
	goose -dir ./migration postgres "host=postgres port=5432 user=postgres password=password dbname=postgres sslmode=disable" up

.PHONY: run-local
run-local:
	go run cmd/video_service/main.go

.PHONY: test
test:
	go test ./...
