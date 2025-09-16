# Server 部分代码学习指南

本指南帮助您理解博客项目中 Go 语言后端的代码结构、工作流程和核心功能。

## 项目结构

server 目录采用典型的 Go Web 应用分层结构，主要包含以下几个部分：

```
server/
├── main.go          # 应用入口
├── config/
│   └── config.go    # 配置管理
├── models/
│   └── blog.go      # 数据模型
├── controllers/
│   └── blog_controller.go  # 控制器
├── routes/
│   └── blog_routes.go      # 路由定义
├── middlewares/
│   └── logger.go    # 中间件
└── utils/
    └── response.go  # 工具函数
```

## 核心组件解析

### 1. 应用入口 (main.go)

`main.go` 是整个后端服务的入口文件，主要负责：

- 初始化配置 (`config.InitConfig()`)
- 连接数据库 (`config.InitDB()`)
- 自动迁移数据库表结构 (`db.AutoMigrate(&models.Blog{})`)
- 添加测试数据
- 创建和配置 Gin 路由器
- 注册中间件
- 注册路由
- 启动 HTTP 服务器

### 2. 配置管理 (config.go)

`config.go` 负责应用配置的管理，主要功能：

- 定义全局配置结构体 `Config`
- 提供 `InitConfig()` 函数从环境变量或默认值初始化配置
- 提供 `InitDB()` 函数创建并返回数据库连接

### 3. 数据模型 (blog.go)

`blog.go` 定义了博客文章的数据模型，使用 GORM ORM 框架：

- 定义 `Blog` 结构体，表示博客文章
- 使用标签定义数据库表结构和 JSON 序列化规则
- 实现 `TableName()` 方法自定义数据库表名

### 4. 控制器 (blog_controller.go)

`blog_controller.go` 包含业务逻辑处理，实现了博客文章的 CRUD 操作：

- 定义 `BlogController` 结构体，持有数据库连接
- 提供 `GetAllBlogs`、`GetBlogByID`、`CreateBlog`、`UpdateBlog` 和 `DeleteBlog` 方法
- 使用 Swagger 注解描述 API 接口
- 处理请求参数、验证、数据库操作和响应

### 5. 路由 (blog_routes.go)

`blog_routes.go` 负责路由注册，将 HTTP 请求映射到对应的控制器方法：

- 定义 `RegisterBlogRoutes` 函数注册所有博客相关路由
- 创建控制器实例
- 定义路由组 `/blogs`
- 注册各 HTTP 方法对应的路由和处理函数

### 6. 工具函数 (response.go)

`response.go` 提供统一的 HTTP 响应格式化工具：

- `SuccessResponse`: 返回成功响应
- `ErrorResponse`: 返回错误响应
- `SuccessListResponse`: 返回带分页的列表响应

### 7. 中间件 (logger.go)

`logger.go` 实现了请求日志记录中间件（目前大部分代码被注释掉），以及 `main.go` 中定义的 CORS 中间件。

## 请求处理流程

一个典型的请求处理流程如下：

1. 客户端发送 HTTP 请求到服务器
2. 请求经过中间件处理（CORS、日志等）
3. Gin 路由器根据 URL 和 HTTP 方法匹配到对应的路由处理函数
4. 控制器方法执行：解析参数、验证、执行数据库操作
5. 使用工具函数返回统一格式的 JSON 响应

## API 接口列表

博客系统提供以下 API 接口：

- `GET /blogs` - 获取所有博客文章（支持分页）
- `GET /blogs/:id` - 获取单篇博客文章
- `POST /blogs` - 创建新的博客文章
- `PUT /blogs/:id` - 更新博客文章
- `DELETE /blogs/:id` - 删除博客文章

## 数据库设计

系统使用 MySQL 数据库，只有一个 `blogs` 表，包含以下字段：

- `id` - 主键，自动递增
- `title` - 博客标题
- `content` - 博客内容（Markdown 格式）
- `author` - 作者
- `created_at` - 创建时间
- `updated_at` - 更新时间

## 学习建议

1. 从 `main.go` 开始，了解应用的整体启动流程
2. 学习 `config.go` 理解配置管理和数据库连接
3. 查看 `models/blog.go` 了解数据模型设计
4. 学习控制器中的各个方法，掌握业务逻辑实现
5. 理解路由注册和中间件的工作原理
6. 尝试修改或扩展现有功能，加深理解

## 扩展建议

- 完善日志中间件功能
- 添加用户认证和授权
- 实现文章分类和标签功能
- 添加评论系统
- 实现文件上传功能（如博客图片）
- 添加缓存机制提高性能