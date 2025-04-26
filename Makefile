test:
	go test . -v

build:
	go build -o bin/$(APP_NAME) main.go

run:
	air

create-migrate:
	migrate create -ext sql -dir migrations -seq $(name)

db-down:
	migrate -path migrations -database postgres://postgres:postgres@localhost/golang-clean?sslmode=disable down

db-up:
	migrate -path migrations -database postgres://postgres:postgres@localhost/golang-clean?sslmode=disable up

db-test-down:
	migrate -path migrations -database postgres://postgres:postgres@localhost/golang-clean-test?sslmode=disable down

db-test-up:
	migrate -path migrations -database postgres://postgres:postgres@localhost/golang-clean-test?sslmode=disable up

install:
	go mod download
