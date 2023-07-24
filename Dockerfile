# 基础镜像
FROM golang:1.20-alpine3.18 AS builder

# 设置工作目录
WORKDIR /app

# 设置代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 复制go.mod和go.sum文件并下载依赖项
COPY go.mod go.sum ./
RUN go mod tidy

# 复制所有文件到工作目录
COPY . .

# 构建应用程序
RUN go build -o myapp

# 使用轻量级基础镜像
FROM alpine:3.14

# 设置工作目录
WORKDIR /app

# 从builder阶段复制构建的应用程序
COPY --from=builder /app/myapp .

# 运行应用程序
CMD ["./myapp"]
