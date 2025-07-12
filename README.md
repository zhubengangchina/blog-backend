# Blog Backend API

这是一个基于Golang的博客后端API系统，提供用户认证、文章管理和评论功能。

## 项目运行环境

- Go 1.24.1
- MySQL 数据库
- Windows/Linux/MacOS 操作系统

## 项目结构

```
blog-backend/
├── config/         # 配置相关代码
├── controllers/    # 控制器，处理HTTP请求
├── middleware/     # 中间件，如JWT认证、日志
├── models/         # 数据模型定义
├── routes/         # 路由配置
├── utils/          # 工具函数
├── go.mod          # Go模块依赖
├── go.sum          # Go模块校验和
└── main.go         # 应用入口
```

## 功能特性

- 用户注册与登录（JWT认证）
- 文章的创建、读取、更新和删除
- 评论管理
- 日志记录
- 数据库自动迁移

## 依赖安装步骤

### 1. 安装Go环境

1. 下载并安装Go 1.24.1或更高版本：
   - Windows: 访问 [Go官方下载页面](https://golang.org/dl/) 下载安装包并运行
   - Linux: `wget https://golang.org/dl/go1.24.1.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz`
   - MacOS: `brew install go` 或访问官方下载页面

2. 配置环境变量：
   - Windows: 系统属性 -> 环境变量 -> 添加 GOPATH 和 GOROOT
   - Linux/MacOS: 在 ~/.bashrc 或 ~/.zshrc 中添加:
     ```
     export GOROOT=/usr/local/go
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
     ```

3. 验证Go安装：
   ```bash
   go version
   ```

### 2. 安装MySQL数据库

1. 下载并安装MySQL：
   - Windows: 从[MySQL官网](https://dev.mysql.com/downloads/installer/)下载安装包
   - Linux: `sudo apt install mysql-server` (Ubuntu) 或 `sudo yum install mysql-server` (CentOS)
   - MacOS: `brew install mysql`

2. 启动MySQL服务：
   - Windows: 通过服务管理或MySQL安装目录的bin文件夹中运行 `mysqld`
   - Linux: `sudo systemctl start mysql`
   - MacOS: `brew services start mysql`

3. 设置MySQL root密码（如果安装过程中未设置）：
   ```bash
   sudo mysql_secure_installation
   ```

### 3. 安装项目依赖

1. 克隆项目到本地：
   ```bash
   git clone <repository-url>
   cd blog-backend
   ```

2. 初始化Go模块（如果是从零开始）：
   ```bash
   go mod init blog-backend
   ```

3. 安装主要依赖：
   ```bash
   # Web框架
   go get -u github.com/gin-gonic/gin@v1.10.1
   
   # ORM框架
   go get -u gorm.io/gorm@v1.30.0
   go get -u gorm.io/driver/mysql@v1.6.0
   
   # JWT认证
   go get -u github.com/golang-jwt/jwt/v5@v5.2.2
   
   # 环境变量加载
   go get -u github.com/joho/godotenv@v1.5.1
   
   # 日志
   go get -u go.uber.org/zap@v1.27.0
   
   # 密码加密
   go get -u golang.org/x/crypto
   ```

4. 同步所有依赖（解决依赖冲突和版本问题）：
   ```bash
   go mod tidy
   ```

5. 验证依赖安装：
   ```bash
   go list -m all
   ```

## 配置环境变量

在项目根目录创建`.env`文件，包含以下配置：

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_db_user
DB_PASS=your_db_password
DB_NAME=blog_db
JWT_SECRET=your_jwt_secret_key
```

## 数据库设置

1. 创建MySQL数据库：

```sql
CREATE DATABASE blog_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 项目启动时会自动创建必要的数据表。

## 启动方式

1. 确保已正确配置环境变量和数据库。

2. 在项目根目录执行：

```bash
go run main.go
```

3. 服务将在`http://localhost:8080`启动。

## API端点

### 认证相关

- `POST /auth/register` - 用户注册
- `POST /auth/login` - 用户登录

### 文章相关

- `GET /posts` - 获取所有文章
- `GET /posts/:id` - 获取特定文章
- `POST /posts` - 创建新文章（需要认证）
- `PUT /posts/:id` - 更新文章（需要认证）
- `DELETE /posts/:id` - 删除文章（需要认证）

### 评论相关

- `GET /posts/:id/comments` - 获取文章评论
- `POST /posts/:id/comments` - 添加评论（需要认证）
- `DELETE /comments/:id` - 删除评论（需要认证）

## 注意事项

- API请求需要在Header中添加`Authorization: Bearer <token>`进行身份验证
- 所有POST和PUT请求的数据格式为JSON
- 错误响应会包含`error`字段描述错误信息 