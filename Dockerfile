FROM golang:latest AS build_base

# RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/task-manager-go

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN CGO_ENABLED=0 go build -o ./out/task-manager-go .

# Start fresh from a smaller image
FROM gcr.io/distroless/static

COPY --from=build_base /tmp/task-manager-go/out/task-manager-go /
COPY --from=build_base /tmp/task-manager-go/config.yml /

# This container exposes port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT [ "/task-manager-go" ]