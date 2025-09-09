# 博客系统前端

这是一个基于Vue 3的博客系统前端实现，支持博客文章的添加、修改、删除和查看功能，并提供Markdown编辑器支持。

## 技术栈

- Vue 3
- Vite
- Vue Router
- Axios
- Markdown-it

## 功能特性

1. **博客列表**：展示所有博客文章，支持删除操作
2. **博客详情**：查看单篇博客文章，支持Markdown渲染
3. **博客创建**：创建新的博客文章，支持Markdown编辑和预览
4. **博客编辑**：编辑现有博客文章，支持Markdown编辑和预览
5. **响应式设计**：适配不同屏幕尺寸

## 项目结构

```
src/
├── views/           # 视图组件
│   ├── BlogList.vue     # 博客列表页
│   ├── BlogDetail.vue   # 博客详情页
│   ├── BlogEdit.vue     # 博客编辑页
│   └── BlogCreate.vue   # 博客创建页
├── router/          # 路由配置
│   └── index.js         # 路由定义
├── services/        # API服务
│   └── api.js           # API接口封装
├── App.vue          # 应用主组件
└── main.js          # 应用入口
```

## 安装和运行

```bash
# 安装依赖
npm install

# 开发模式运行
npm run dev

# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

## 后端API说明

请参考项目根目录下的 `API_DOC.md` 文件，了解所需的RESTful API接口定义。

## 注意事项

1. 本项目使用模拟数据进行展示，实际使用时需要连接真实的后端API
2. 修改 `src/services/api.js` 中的 `baseURL` 以指向实际的后端服务地址
3. 如需添加认证功能，请在API拦截器中实现token的添加和验证逻辑
