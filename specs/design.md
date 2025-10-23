# 技术方案设计

## 系统架构

### 整体架构
```
┌─────────────────┐    HTTP/HTTPS    ┌─────────────────┐
│   前端 (Vue.js)  │ ◄──────────────► │   后端 (Go)     │
│   Port: 5173    │                  │   Port: 8080    │
└─────────────────┘                  └─────────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │   数据库 (SQLite)│
                                     │   文件存储      │
                                     └─────────────────┘
```

### 技术栈选型

#### 前端技术栈
- **框架**: Vue.js 3.x - 现代化响应式框架
- **构建工具**: Vite - 快速构建和热重载
- **UI组件**: 原生HTML/CSS - 保持轻量级
- **路由**: Vue Router - 单页面应用路由
- **HTTP客户端**: Axios - API请求处理
- **Markdown渲染**: marked.js - Markdown解析和渲染
- **图表渲染**: mermaid.js - 流程图、时序图等图表渲染
- **代码高亮**: highlight.js - 多语言代码语法高亮
- **工具库**: @vueuse/core - Vue组合式API工具库
- **编辑器**: 掘金风格AdvancedEditor - 支持实时预览、分屏编辑、工具栏操作

#### 后端技术栈
- **语言**: Go 1.19+ - 高性能并发语言
- **Web框架**: Gin - 轻量级HTTP框架
- **数据库**: SQLite - 嵌入式数据库，无需额外配置
- **认证**: JWT - 无状态认证机制
- **中间件**: 自定义中间件 - 认证、日志、CORS

#### 数据库设计

**用户表 (users)**
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

**博客表 (blogs)**
```sql
CREATE TABLE blogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id)
);
```

### API接口设计

#### 用户认证接口
```
POST /api/auth/register - 用户注册
POST /api/auth/login    - 用户登录
POST /api/auth/logout   - 用户登出
GET  /api/auth/profile  - 获取用户信息
```

#### 博客管理接口
```
GET    /api/blogs        - 获取博客列表
GET    /api/blogs/:id    - 获取博客详情
POST   /api/blogs        - 创建博客
PUT    /api/blogs/:id    - 更新博客
DELETE /api/blogs/:id    - 删除博客
```

#### 文件上传接口
```
POST   /api/upload/image - 图片上传接口
GET    /uploads/*         - 静态文件访问
```

### 安全策略

#### 认证机制
- **JWT Token**: 无状态认证，包含用户ID和过期时间
- **密码加密**: 使用bcrypt进行密码哈希
- **Token过期**: 设置合理的过期时间（24小时）

#### 数据验证
- **输入验证**: 前后端双重验证
- **SQL注入防护**: 使用参数化查询
- **XSS防护**: 输出转义和CSP策略

#### CORS配置
- **允许来源**: 开发环境允许localhost，生产环境限制域名
- **允许方法**: GET, POST, PUT, DELETE
- **允许头部**: Content-Type, Authorization

### 测试策略

#### 单元测试
- **后端**: 使用Go内置testing包
- **前端**: 使用Vue Test Utils + Jest
- **覆盖率**: 目标80%以上

#### 集成测试
- **API测试**: 使用Postman或自动化测试工具
- **端到端测试**: 关键用户流程测试

### 部署方案

#### 开发环境
- **前端**: Vite开发服务器 (localhost:5173)
- **后端**: Go直接运行 (localhost:8080)
- **数据库**: SQLite文件存储

#### 生产环境
- **前端**: 静态文件部署到CDN或Web服务器
- **后端**: 编译为二进制文件部署到服务器
- **数据库**: 考虑升级到PostgreSQL或MySQL
- **反向代理**: 使用Nginx处理静态文件和API路由

### 性能优化

#### 前端优化
- **代码分割**: 按路由分割代码
- **懒加载**: 图片和组件懒加载
- **缓存策略**: 静态资源缓存

#### 后端优化
- **连接池**: 数据库连接池管理
- **缓存**: 热点数据缓存
- **压缩**: Gzip压缩响应内容

### 监控和日志

#### 日志记录
- **访问日志**: 记录所有HTTP请求
- **错误日志**: 记录系统错误和异常
- **业务日志**: 记录关键业务操作

#### 监控指标
- **响应时间**: API响应时间监控
- **错误率**: 系统错误率统计
- **并发数**: 同时在线用户数
