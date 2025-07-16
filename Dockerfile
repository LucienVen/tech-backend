# ---------- Build stage ----------
FROM golang:1.24.1-bookworm AS builder

# 设置 Go 模块代理（可选）
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

WORKDIR /app

# 先缓存 go.mod/go.sum
COPY go.mod go.sum ./
RUN go mod download

# 再复制源码
COPY . .

# 注意这里我们不强制 CGO_ENABLED=0
# 保证 cgo 能工作
RUN go build -o tech-backend ./cmd/main.go


# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/tech-backend .

# 暴露端口
EXPOSE 8080

# 启动命令
ENTRYPOINT ["/app/tech-backend"]
