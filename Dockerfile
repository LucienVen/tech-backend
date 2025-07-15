FROM golang:1.24.1 AS builder

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tech-backend ./cmd/main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y tzdata \
  && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone \
  && dpkg-reconfigure -f noninteractive tzdata \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/tech-backend .

EXPOSE 8080
CMD ["./tech-backend"]
