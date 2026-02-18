FROM golang:1.25.4-alpine

RUN apk add --no-cache git build-base

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080 2345
