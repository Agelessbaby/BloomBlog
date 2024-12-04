# 第一阶段：编译阶段
FROM golang:1.23.2 AS builder

# 设置工作目录
WORKDIR /app

# 将源码添加到工作目录
COPY . .

RUN go mod tidy

RUN chmod +x ./cmd/publish/build.sh

# 编译 Go 应用
RUN cd ./cmd/publish/&&./build.sh

# 第二阶段：运行阶段
FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 从第一阶段复制编译好的二进制文件
COPY --from=builder /app/cmd/publish/output/bin/publish .

# 复制配置文件和模板目录
COPY --from=builder /app/config ./config/

# 确保二进制可执行
RUN chmod +x ./publish

# 暴露端口
EXPOSE 8084

# 运行二进制文件
ENTRYPOINT ["./publish"]

