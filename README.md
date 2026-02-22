# jfeng_blog 个人博客系统

一个使用 Go 语言开发的轻量级个人博客系统，基于 Gin 框架构建。

## 技术栈

| 类别 | 技术 |
|------|------|
| 后端框架 | Gin v1.11.0 |
| ORM | GORM v1.31.1 |
| 数据库 | MySQL |
| 配置管理 | Viper |
| 日志 | Zap |
| API 文档 | Swagger + gin-swagger |
| 热重载 | Air |

## 项目结构

```
jfeng_blog/
├── api/                      # API 层
│   ├── docs/                 # Swagger 文档
│   └── dto/                  # 数据传输对象
│       ├── response/         # 响应结构
│       └── userResponse.go   # 用户相关响应
├── cmd/
│   └── server/               # 程序入口
│       └── apiServer.go
├── config/                   # 配置文件
│   ├── config.yaml          # 配置文件
│   └── config.example.yaml  # 配置示例
├── databases/                # 数据库脚本
│   └── jfeng_blog_20260213.sql
├── docs/                     # 文档
│   ├── design.md            # 设计文档
│   └── develop.md          # 开发文档
├── internal/                 # 内部包
│   ├── handler/             # 处理器
│   ├── middleware/         # 中间件
│   ├── model/               # 数据模型
│   ├── repository/         # 数据访问层
│   ├── router/              # 路由
│   │   └── api_v1/         # API v1 版本
│   ├── service/             # 业务逻辑层
│   └── utils/               # 工具函数
├── .air.toml               # Air 热重载配置
├── go.mod                  # Go 模块依赖
└── README.md
```

## 快速开始

### 环境要求

- Go 1.25+
- MySQL 5.7+
- Git

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/your-repo/jfeng_blog.git
cd jfeng_blog
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置数据库

复制配置示例文件并修改：
```bash
cp config/config.example.yaml config/config.yaml
```

编辑 `config/config.yaml`，配置 MySQL 连接信息。

4. 初始化数据库

执行数据库脚本创建表：
```bash
mysql -u root -p < databases/jfeng_blog_20260213.sql
```

5. 运行项目

开发模式（使用 air 热重载）：
```bash
air
```

或者直接运行：
```bash
go run cmd/server/apiServer.go
```

### 访问服务

- 服务地址：http://localhost:8080
- Swagger 文档：http://localhost:8080/docs

## API 文档

详细的 API 文档请访问 Swagger UI：http://localhost:8080/docs

## 功能特性

- [x] 用户管理
- [ ] 文章管理
- [ ] 分类管理
- [ ] 标签管理
- [ ] 评论管理
- [ ] 站点设置

## 开发指南

详细开发指南请参考 [开发文档](docs/develop.md)

## 设计文档

详细设计文档请参考 [设计文档](docs/design.md)

## 许可证

MIT License
