# 使用基于 Fedora 的 Buildah 镜像
FROM quay.io/buildah/upstream:latest

# 安装必要的依赖工具
RUN dnf install -y \
    curl \
    tar \
    && dnf clean all

# 安装 doctl (DigitalOcean CLI)
RUN curl -fsSL https://github.com/digitalocean/doctl/releases/download/v1.100.1/doctl-1.100.1-linux-amd64.tar.gz | tar -xz -C /usr/local/bin

# 验证安装
RUN buildah --version && doctl version

# 设置默认工作目录
WORKDIR /workspace
