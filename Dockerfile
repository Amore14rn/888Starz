FROM golang:1.20-alpine AS builder

# Определение метаданных для образа
LABEL service="888Starz"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pkg ./pkg
COPY internal ./internal
COPY cmd ./cmd

RUN go build -o /app/888Starz ./cmd/

EXPOSE 8080

CMD ["/app/888Starz"]