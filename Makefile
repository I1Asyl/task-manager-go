$(shell touch .env)
include .env


run:
	go run main.go

build:
	go mod tidy && go build -o bin/

database-up:
	cd database/migrations && go run main.go up $(DSN)

database-down:
	cd database/migrations && go run main.go down $(DSN)

docker-build-run:
	docker build -t task-manager-go:0.1 .;
	docker run --name=task-manager-go --rm -p 8080:8080 task-manager-go:0.1 

docker-run:
	docker run --name=task-manager-go --rm -p 8080:8080 task-manager-go:0.1 

test-services:
	cd pkg/services && go test -v -cover