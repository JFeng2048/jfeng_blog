# 设计文档

## 系统架构

### 整体架构

本项目采用经典的 **三层架构**（或四层架构）设计：

```
┌─────────────────────────────────────────┐
│              Router 层                  │  路由层
├─────────────────────────────────────────┤
│             Handler 层                  │  处理器层 (接收请求)
├─────────────────────────────────────────┤
│             Service 层                  │  业务逻辑层
├─────────────────────────────────────────┤
│           Repository 层                 │  数据访问层
├─────────────────────────────────────────┤
│              Model 层                   │  数据模型层
└─────────────────────────────────────────┘
                     ↓
┌─────────────────────────────────────────┐
│               Database                   │  MySQL
└─────────────────────────────────────────┘
```

### 分层说明

| 层级 | 职责 | 目录 |
|------|------|------|
| Router | 路由注册、请求分发 | internal/router |
| Handler | 请求参数校验、调用 Service、返回响应 | internal/handler |
| Service | 业务逻辑处理 | internal/service |
| Repository | 数据库 CRUD 操作 | internal/repository |
| Model | 数据结构定义 | internal/model |

## 数据库设计

### 数据表结构

#### 1. users - 用户表

| 字段 | 类型 | 说明 |
|------|------|------|
| user_id | int | 用户ID（主键） |
| username | varchar(255) | 用户名（唯一） |
| display_name | varchar(255) | 显示名称 |
| avatar_url | varchar(255) | 头像URL |
| bio | text | 个人简介 |
| role | varchar(255) | 角色（admin/author/guest） |
| email | varchar(255) | 邮箱（唯一） |
| password_hash | varchar(255) | 密码哈希 |
| social_links | json | 社交链接 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

#### 2. post - 文章表

| 字段 | 类型 | 说明 |
|------|------|------|
| post_id | int | 文章ID（主键） |
| title | varchar(255) | 文章标题 |
| slug | varchar(255) | 文章别名（唯一） |
| content | longtext | 文章内容 |
| excerpt | text | 文章摘要 |
| cover_image | varchar(255) | 封面图 |
| status | varchar(255) | 状态（draft/published） |
| view_count | int | 浏览次数 |
| is_featured | tinyint | 是否推荐 |
| is_sticky | tinyint | 是否置顶 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

#### 3. project - 项目表

| 字段 | 类型 | 说明 |
|------|------|------|
| project_id | int | 项目ID（主键） |
| name | varchar(255) | 项目名称 |
| description | text | 项目描述 |
| url | varchar(255) | 项目链接 |
| icon | varchar(255) | 项目图标 |
| order_num | int | 排序 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

#### 4. sys_tag - 标签表

| 字段 | 类型 | 说明 |
|------|------|------|
| tag_id | int | 标签ID（主键） |
| name | varchar(255) | 标签名称 |
| slug | varchar(255) | 标签别名 |
| description | text | 标签描述 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

#### 5. timeline - 时间线表

| 字段 | 类型 | 说明 |
|------|------|------|
| timeline_id | int | 时间线ID（主键） |
| title | varchar(255) | 标题 |
| content | text | 内容 |
| date | date | 日期 |
| type | varchar(255) | 类型 |
| order_num | int | 排序 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

## API 设计

### 路由规范

- 基础路径：`/api/v1`
- RESTful 风格
- 使用名词而非动词（如 `/users` 而非 `/getUsers`）

### 响应格式

成功响应：
```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

错误响应：
```json
{
  "code": 400,
  "msg": "error message",
  "data": null
}
```

### 已实现 API

#### 用户模块

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/user/:id | 获取用户信息 |

## 配置说明

配置文件位于 `config/config.yaml`：

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  name: "jfeng_blog"

system:
  version: "1.0.0"

database:
  host: "localhost"
  port: 3306
  username: "root"
  password: ""
  database: "jfeng_blog"
  max_idle_conns: 10
  max_open_conns: 25
  conn_max_lifetime: 300

mysql:
  host: "localhost"
  port: 3306
  username: "root"
  password: ""
  database: "jfeng_blog"
```

## 中间件

项目使用以下中间件：

1. **Logger** - 请求日志记录
2. **Recovery** - 异常恢复
3. **CORS** - 跨域资源共享

## 技术选型理由

### Gin 框架
- 性能高
- 路由设计简洁
- 中间件生态丰富

### GORM
- Go 语言最流行的 ORM
- 支持链式操作
- 自动迁移

### Zap 日志
- Uber 出品，性能优异
- 结构化日志
- 支持多种输出格式

### Viper
- 支持多种配置格式
- 环境变量支持
- 配置热加载
