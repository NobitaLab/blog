# 博客系统项目

这是一个现代化的博客管理系统，采用前后端分离架构，提供完整的博客内容管理功能。

## 项目架构

### 技术栈
- **前端**: Vue.js 3 + Vite + Vue Router + Axios
- **后端**: Go + Gin + SQLite + JWT
- **数据库**: SQLite (开发环境)
- **构建工具**: Vite (前端) + Go Build (后端)

### 项目结构
```
blog/
├── client/                 # 前端Vue.js应用
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   ├── services/      # API服务
│   │   └── router/        # 路由配置
│   ├── public/            # 静态资源
│   └── dist/              # 构建输出
├── server/                # 后端Go API服务
│   ├── controllers/       # 控制器
│   ├── models/           # 数据模型
│   ├── middlewares/      # 中间件
│   ├── routes/           # 路由定义
│   └── utils/            # 工具函数
├── specs/                # 项目规范文档
│   ├── requirements.md   # 需求文档
│   ├── design.md         # 技术方案设计
│   └── tasks.md          # 任务拆分
└── .cursor/rules/        # 开发规则文件
```

## 核心功能

### 用户认证
- 用户注册和登录
- JWT Token认证
- 密码加密存储
- 会话管理

### 博客管理
- 博客文章的CRUD操作
- Markdown格式支持
- 实时预览功能
- 文章搜索和分页

### 界面设计
- 响应式设计，支持多设备
- 现代化UI界面
- 暗色主题支持
- 无障碍访问支持

## 开发环境配置

### 环境要求
- Go 1.19+ (使用GOPATH模式)
- Node.js 16+
- npm 或 yarn

### 启动步骤

1. **启动后端服务**
   ```bash
   cd server
   go run main.go
   ```
   后端服务将在 http://localhost:8080 启动

2. **启动前端服务**
   ```bash
   cd client
   npm install
   npm run dev
   ```
   前端服务将在 http://localhost:5173 启动

3. **访问应用**
   - 前端地址: http://localhost:5173
   - API文档: http://localhost:8080/api/docs

## API接口

### 用户认证
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/logout` - 用户登出
- `GET /api/auth/profile` - 获取用户信息

### 博客管理
- `GET /api/blogs` - 获取博客列表
- `GET /api/blogs/:id` - 获取博客详情
- `POST /api/blogs` - 创建博客
- `PUT /api/blogs/:id` - 更新博客
- `DELETE /api/blogs/:id` - 删除博客

## 开发规范

### 代码规范
- 遵循Go官方代码规范
- 使用ESLint进行前端代码检查
- 统一的代码格式化配置

### 项目规范
- 使用Git进行版本控制
- 遵循语义化版本规范
- 完善的代码注释和文档

### 测试规范
- 后端单元测试覆盖
- 前端组件测试
- API接口测试

## 部署说明

### 开发环境
- 使用SQLite数据库
- 热重载开发模式
- 本地调试配置

### 生产环境
- 可升级到PostgreSQL/MySQL
- 静态资源CDN加速
- 反向代理配置
- 监控和日志系统

## 维护指南

### 数据库维护
- 定期备份数据库
- 监控数据库性能
- 优化查询语句

### 代码维护
- 定期更新依赖包
- 代码重构和优化
- 安全漏洞修复

### 文档维护
- 更新API文档
- 完善用户手册
- 维护开发文档

## 贡献指南

1. Fork项目仓库
2. 创建功能分支
3. 提交代码变更
4. 创建Pull Request
5. 代码审查和合并

## 许可证

本项目采用MIT许可证，详情请查看LICENSE文件。