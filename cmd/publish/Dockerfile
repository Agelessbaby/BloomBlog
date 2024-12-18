# 第一阶段：编译阶段
FROM golang:1.23.2 AS builder

# 设置工作目录
WORKDIR /app

# 将源码添加到工作目录
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# 管理依赖
RUN go mod tidy

# 确保脚本可执行
RUN chmod +x ./cmd/publish/build.sh

# 设置目标平台

# 编译 Go 应用
RUN cd ./cmd/publish/ && ./build.sh

# 第二阶段：运行阶段
# 明确指定稳定的 Ubuntu 版本
FROM ubuntu:22.04

# 更新包管理器并安装必需的包
RUN apt-get update --fix-missing && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*


# 设置工作目录
WORKDIR /app

# 从第一阶段复制编译好的二进制文件
COPY --from=builder /app/cmd/publish/output/bin/publish .

# 复制配置文件和模板目录
COPY --from=builder /app/config ./config/

# 确保二进制文件可执行
RUN chmod +x ./publish

# 暴露端口
EXPOSE 8084

# 设置容器的入口点
ENTRYPOINT ["./publish"]
