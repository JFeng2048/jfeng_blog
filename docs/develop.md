# 开发文档

## 开发环境搭建

### 环境要求

- Go 1.25+
- MySQL 5.7+
- Git
- VS Code 或 GoLand（推荐）

### 安装 Go

1. 下载 Go 安装包：https://go.dev/dl/
2. 安装并配置环境变量
3. 验证安装：
```bash
go version
```

### 安装 MySQL

1. 下载 MySQL：https://dev.mysql.com/downloads/mysql/
2. 安装并启动服务
3. 创建数据库：
```bash
mysql -u root -p
CREATE DATABASE jfeng_blog;
```

### 克隆项目

```bash
git clone <repository-url>
cd jfeng_blog
```

### 安装依赖

```bash
go mod tidy
```

### 配置项目

1. 复制配置示例：
```bash
cp config/config.example.yaml config/config.yaml
```

2. 编辑 `config/config.yaml`，修改数据库连接信息：
```yaml
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "your-password"
  database: "jfeng_blog"
```

### 初始化数据库

```bash
mysql -u root -p jfeng_blog < databases/jfeng_blog_20260213.sql
```

### 运行项目

#### 方式一：使用 Air 热重载（推荐）

```bash
# 安装 air
go install github.com/air-verse/air@latest

# 运行
air
```

#### 方式二：直接运行

```bash
go run cmd/server/apiServer.go
```

### 访问服务

- 服务地址：http://localhost:8080
- Swagger 文档：http://localhost:8080/docs

---

## 代码规范

### 项目目录规范

```
internal/
├── handler/          # 处理请求和响应
├── middleware/       # 中间件
├── model/            # 数据模型
├── repository/      # 数据访问层
├── router/          # 路由注册
│   └── api_v1/     # API v1 版本
├── service/         # 业务逻辑
└── utils/           # 工具函数
```

### 命名规范

- **文件命名**：使用下划线命名法，如 `user.go`、`user_service.go`
- **结构体命名**：使用大驼峰命名法，如 `User`、`UserService`
- **函数命名**：使用大驼峰命名法，如 `GetUserByID`
- **变量命名**：使用小驼峰命名法，如 `userName`、`userList`
- **常量命名**：使用全大写下划线，如 `MAX_COUNT`、`DEFAULT_NAME`

### 分层规范

1. **Handler 层**：只负责接收请求、参数校验、调用 Service、返回响应
2. **Service 层**：处理业务逻辑，不直接操作数据库
3. **Repository 层**：只负责数据库 CRUD 操作
4. **Model 层**：定义数据结构

### 注释规范

- 所有导出函数必须添加中文注释
- 注释应简洁说明函数功能，不说明参数和返回值

---

## 常用命令

### Go 命令

```bash
# 运行项目
go run cmd/server/apiServer.go

# 编译项目
go build -o ./tmp/main.exe ./cmd/server

# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 代码格式检查
go fmt ./...

# 代码检查
go vet ./...
```

### Air 命令

```bash
# 运行 air（开发模式）
air

# 查看 air 配置帮助
air -h
```

### Swagger 命令

```bash
# 生成 Swagger 文档
swag init -g ./cmd/server/apiServer.go -o ./api/docs
```

---

## 添加新功能的步骤

### 1. 创建数据模型

在 `internal/model/` 下创建新的模型文件，如 `post.go`：

```go
package model

import "time"

type Post struct {
    ID        int       `gorm:"column:post_id;primaryKey"`
    Title     string    `gorm:"column:title;type:varchar(255)"`
    Content   string    `gorm:"column:content;type:longtext"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Post) TableName() string {
    return "post"
}
```

### 2. 创建 Repository 层

在 `internal/repository/` 下创建 `post.go`：

```go
package repository

import "jfeng_blog/internal/model"

func GetPostByID(id uint) (*model.Post, error) {
    var post model.Post
    if err := db.Where("post_id = ?", id).First(&post).Error; err != nil {
        return nil, err
    }
    return &post, nil
}
```

### 3. 创建 Service 层

在 `internal/service/post/` 下创建 `post.go`：

```go
package post

import (
    "errors"
    "jfeng_blog/internal/model"
    "jfeng_blog/internal/repository"
    "go.uber.org/zap"
)

func GetPostInfoByID(id int) (*model.Post, error) {
    if id <= 0 {
        return nil, errors.New("invalid post id")
    }
    
    post, err := repository.GetPostByID(uint(id))
    if err != nil {
        zap.L().Error("get post failed", zap.Int("id", id), zap.Error(err))
        return nil, err
    }
    
    return post, nil
}
```

### 4. 创建 Handler 层

在 `internal/handler/` 下创建 `post.go`：

```go
package handler

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "jfeng_blog/api/dto/response"
    "jfeng_blog/internal/service/post"
)

// @Summary 根据ID获取文章
// @Tags 文章
// @Param id path int true "文章ID"
// @Success 200 {object} response.SuccessResponse
// @Router /post/{id} [get]
func GetPostInfoByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(400, response.ErrorResponse{
            Code: 400,
            Msg:  "invalid id",
        })
        return
    }
    
    p, err := post.GetPostInfoByID(id)
    if err != nil {
        c.JSON(500, response.ErrorResponse{
            Code: 500,
            Msg:  "get post failed",
        })
        return
    }
    
    c.JSON(200, response.SuccessResponse{
        Code: 200,
        Msg:  "success",
        Data: p,
    })
}
```

### 5. 注册路由

在 `internal/router/api_v1/` 下找到或创建 `post.go`：

```go
package api_v1

import (
    "github.com/gin-gonic/gin"
    "jfeng_blog/internal/handler"
)

func RegisterPostRoutes(r *gin.RouterGroup) {
    postGroup := r.Group("/post")
    {
        postGroup.GET("/:id", handler.GetPostInfoByID)
    }
}
```

在 `internal/router/api_v1/register.go` 中注册路由：

```go
func RegisterV1(r *gin.RouterGroup) {
    RegisterUserRoutes(r)
    RegisterPostRoutes(r)  // 添加这行
}
```

### 6. 生成 Swagger 文档

```bash
swag init -g ./cmd/server/apiServer.go -o ./api/docs
```

---

## 常见问题

### 1. 数据库连接失败

检查 `config/config.yaml` 中的数据库配置是否正确，确保 MySQL 服务已启动。

### 2. Swagger 文档不更新

运行 `swag init` 重新生成文档：

```bash
swag init -g ./cmd/server/apiServer.go -o ./api/docs
```

### 3. Air 热重载循环

如果 `api/docs` 目录被监听，会导致热重载循环。在 `.air.toml` 中添加排除目录：

```toml
exclude_dir = ["api/docs", "tmp", "vendor"]
```

### 4. 时间字段解析失败

确保 DSN 中包含 `parseTime=True`：

```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", ...)
```

---

## 开发工具推荐

- **IDE**：VS Code、GoLand
- **API 测试**：Postman、Apifox
- **数据库客户端**：MySQL Workbench、Navicat、DBeaver
- **Git 客户端**：Git Bash、GitHub Desktop
