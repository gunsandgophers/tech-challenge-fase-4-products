FROM golang:1.22-alpine3.20

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

# Download Go modules
COPY . .

RUN go mod download

RUN go build -o build/main

ENTRYPOINT ./build/main

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080
