FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 安装依赖
RUN go mod tidy

# 构建应用
RUN go build -o api ./cmd/api


# 第二阶段：运行阶段
FROM ubuntu:latest

# 设置工作目录
WORKDIR /app

# 从第一阶段复制编译好的二进制文件
COPY --from=builder /app/api .

# 复制配置文件和模板目录
COPY --from=builder /app/config ./config/

COPY --from=builder /app/docs ./docs

# 确保二进制可执行
RUN chmod +x ./api

# 暴露端口
EXPOSE 8088

# 运行二进制文件
ENTRYPOINT ["./api"]

