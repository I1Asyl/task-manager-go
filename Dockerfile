FROM golang:1.21 AS build_base

# RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/task-manager-go

# We want to populate the module cache based on the go.{mod,sum} files. update#2
COPY . .

RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 go build -o ./out/task-manager-go .

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["/tmp/task-manager-go/out/task-manager-go"]
