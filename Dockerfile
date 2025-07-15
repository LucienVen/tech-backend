FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tech-backend ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/tech-backend .

EXPOSE 8080
CMD ["./tech-backend"]
