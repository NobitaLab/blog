# 博客项目

这是一个简单的博客系统，包含前端和后端两部分。

## 项目结构
- `client/`: 前端Vue.js应用，使用Vite构建
- `server/`: 后端Go语言API服务

## 功能特点
- 博客文章的创建、编辑、查看和删除
- Markdown格式支持
- 响应式设计

## 使用方法
1. 启动后端服务
   ```
   cd server
   go run main.go
   ```

2. 启动前端服务
   ```
   cd client
   npm install
   npm run dev
   ```

3. 访问 http://localhost:5173 查看应用