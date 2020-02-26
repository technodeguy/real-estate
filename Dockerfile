FROM golang:1.13-alpine

RUN mkdir -p /app
WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod /app
COPY go.sum /app

RUN go mod download 

COPY . /app

CMD ["go", "run", "main.go"]