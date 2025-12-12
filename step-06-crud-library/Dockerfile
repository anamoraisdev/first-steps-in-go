# build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server main.go

# final stage
FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/server .

ENV APP_PORT=8080
EXPOSE 8080

CMD ["./server"]
