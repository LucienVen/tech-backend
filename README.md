# tech-backend



# 部署备注

1. 依赖外部 MySQL、PostgreSQL、Redis（这些服务已部署，无需在 compose 里启动，只需配置连接）
2. 系统版本"Debian GNU/Linux 12 (bookworm)"，已安装 Docker & docker-compose。
3. 服务监听端口通过 .env 文件配置。
4. 镜像无需手动推送，代码 push 到 GitHub 后，由 GitHub Actions 自动构建、推送镜像并部署到远程服务器 Docker。



## 详细实现方案

### 1. Dockerfile

- 多阶段构建，最终镜像精简。

- 读取 .env 配置端口。

### 2. docker-compose.yml

- 只定义 tech-backend 服务，依赖外部数据库和 Redis，通过环境变量连接。

### 3. .env 示例

- 提供连接外部 MySQL、PG、Redis 的环境变量，以及服务端口。

### 4. GitHub Actions workflow

- 触发条件：push 到 main 分支

- 步骤：
  - 构建 Docker 镜像
  - 推送到远程服务器（可用 ssh/scp 或远程 docker context）
  - 在远程服务器上 docker-compose up -d