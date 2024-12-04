# 第一阶段：编译阶段
FROM golang:1.23.2 AS builder

# 设置工作目录
WORKDIR /app

# 将源码添加到工作目录
COPY . .

RUN go mod tidy

# 编译 Go 应用
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

