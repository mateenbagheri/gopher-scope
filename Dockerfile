FROM golang:1.25.4-alpine

RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080 40000

CMD ["air", "-c", ".air.toml"]
