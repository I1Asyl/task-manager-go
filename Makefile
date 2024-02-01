run:
	go run main.go

build:
	go mod tidy && go build -o bin/

docker-build-run:
	docker build -t task-manager-go:0.1 .;
	docker run --name=task-manager-go --rm -p 8080:8080 task-manager-go:0.1 

docker-run:
	docker run --name=task-manager-go --rm -p 8080:8080 task-manager-go:0.1 