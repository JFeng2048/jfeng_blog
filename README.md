# jfeng_blog
个人博客网站后端API项目

## 项目简介

这是一个基于Go语言开发的个人博客后端API系统，提供RESTful接口用于展示博客信息、项目信息、个人信息等内容。

## 功能特性

- 博客信息管理：获取所有博客文章和单篇博客详情
- 项目信息展示：获取所有项目和单个项目详情
- 个人信息展示：获取个人简介、技能、教育背景和工作经历

## 技术栈

- Go 语言
- 标准库 net/http
- RESTful API 设计

## API 接口

### 根路径
```
GET /
```
返回API欢迎信息和所有可用端点列表

### 博客接口
```
GET /api/blogs          # 获取所有博客文章
GET /api/blogs/{id}     # 获取指定ID的博客文章
```

### 项目接口
```
GET /api/projects       # 获取所有项目
GET /api/projects/{id}  # 获取指定ID的项目
```

### 个人信息接口
```
GET /api/personal       # 获取个人信息
```

## 快速开始

### 安装依赖
```bash
go mod tidy
```

### 构建项目
```bash
go build -o jfeng_blog main.go
```

### 运行服务器
```bash
./jfeng_blog
```

服务器将在 `http://localhost:8080` 启动

### 测试API

使用curl测试：
```bash
# 测试根路径
curl http://localhost:8080/

# 获取所有博客
curl http://localhost:8080/api/blogs

# 获取指定博客
curl http://localhost:8080/api/blogs/1

# 获取所有项目
curl http://localhost:8080/api/projects

# 获取指定项目
curl http://localhost:8080/api/projects/1

# 获取个人信息
curl http://localhost:8080/api/personal
```

## 项目结构

```
jfeng_blog/
├── main.go              # 主程序入口和路由配置
├── models/              # 数据模型定义
│   ├── blog.go         # 博客模型
│   ├── project.go      # 项目模型
│   └── personal.go     # 个人信息模型
├── handlers/            # HTTP请求处理器
│   ├── blogs.go        # 博客相关处理器
│   ├── projects.go     # 项目相关处理器
│   └── personal.go     # 个人信息处理器
├── data/               # 示例数据
│   ├── blogs.go        # 博客示例数据
│   ├── projects.go     # 项目示例数据
│   └── personal.go     # 个人信息示例数据
├── go.mod              # Go模块依赖
└── README.md           # 项目说明文档
```

## 示例响应

### 博客列表响应示例
```json
[
  {
    "id": 1,
    "title": "开始我的个人博客之旅",
    "content": "这是我的第一篇博客文章...",
    "author": "JFeng",
    "tags": ["博客", "个人成长", "技术"],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z",
    "view_count": 120,
    "category": "随笔",
    "description": "记录开始写博客的初衷和计划"
  }
]
```

### 项目列表响应示例
```json
[
  {
    "id": 1,
    "name": "个人博客系统",
    "description": "基于Go语言开发的个人博客后端API系统",
    "technologies": ["Go", "RESTful API", "JSON"],
    "github_url": "https://github.com/JFeng2048/jfeng_blog",
    "live_url": "",
    "image_url": "",
    "start_date": "2024-01-01T00:00:00Z",
    "status": "active"
  }
]
```

### 个人信息响应示例
```json
{
  "name": "JFeng",
  "title": "全栈开发工程师",
  "bio": "热爱编程，专注于Web开发和软件工程",
  "email": "jfeng@example.com",
  "location": "中国",
  "skills": ["Go", "JavaScript", "Python", "React"],
  "education": [...],
  "experience": [...]
}
```

## 开发说明

当前版本使用内存数据存储（示例数据），后续可以扩展为：
- 连接数据库（MySQL、PostgreSQL等）
- 添加CRUD操作（创建、更新、删除）
- 添加用户认证和授权
- 添加分页和搜索功能
- 添加CORS支持用于前端集成

## 许可证

MIT License
