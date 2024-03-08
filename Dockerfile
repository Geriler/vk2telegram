# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:latest

ENV CGO_ENABLED=1

RUN apt-get update && apt-get upgrade && \
apt-get install --no-install-recommends --assume-yes build-essential libsqlite3-dev git bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/vk2telegram

EXPOSE 80

CMD ["./main"]