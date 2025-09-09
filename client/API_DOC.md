# 博客系统后端API文档

本文档描述了博客系统前端所需的后端RESTful API接口规范。

## 基础URL

所有API请求的基础URL为：`/api`

## API版本

当前版本：v1

## 认证

本API暂未实现认证机制，实际项目中可根据需要添加JWT或其他认证方式。

## 接口列表

### 博客文章相关接口

#### 1. 获取所有博客文章

- **URL**: `/api/blogs`
- **Method**: `GET`
- **Description**: 获取所有博客文章列表
- **Query Parameters**:
  - `page` (可选): 页码，默认为1
  - `limit` (可选): 每页数量，默认为10
- **Success Response**:
  - **Code**: 200
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "data": [
        {
          "id": 1,
          "title": "博客标题",
          "content": "博客内容（Markdown格式）",
          "author": "作者名称",
          "createdAt": "2024-01-15T10:00:00Z",
          "updatedAt": "2024-01-15T10:00:00Z"
        },
        // 更多博客文章...
      ],
      "total": 100,
      "page": 1,
      "limit": 10
    }
    ```
- **Error Response**:
  - **Code**: 500
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "获取博客列表失败"
    }
    ```

#### 2. 获取单篇博客文章

- **URL**: `/api/blogs/:id`
- **Method**: `GET`
- **Description**: 获取指定ID的博客文章详情
- **URL Parameters**:
  - `id`: 博客文章ID
- **Success Response**:
  - **Code**: 200
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "data": {
        "id": 1,
        "title": "博客标题",
        "content": "博客内容（Markdown格式）",
        "author": "作者名称",
        "createdAt": "2024-01-15T10:00:00Z",
        "updatedAt": "2024-01-15T10:00:00Z"
      }
    }
    ```
- **Error Response**:
  - **Code**: 404
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "博客不存在"
    }
    ```
  - **Code**: 500
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "获取博客详情失败"
    }
    ```

#### 3. 创建新的博客文章

- **URL**: `/api/blogs`
- **Method**: `POST`
- **Description**: 创建一篇新的博客文章
- **Request Body**:
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "title": "博客标题",
      "content": "博客内容（Markdown格式）",
      "author": "作者名称"
    }
    ```
- **Success Response**:
  - **Code**: 201
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "data": {
        "id": 3,
        "title": "博客标题",
        "content": "博客内容（Markdown格式）",
        "author": "作者名称",
        "createdAt": "2024-01-20T14:30:00Z",
        "updatedAt": "2024-01-20T14:30:00Z"
      }
    }
    ```
- **Error Response**:
  - **Code**: 400
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "标题和内容不能为空"
    }
    ```
  - **Code**: 500
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "创建博客失败"
    }
    ```

#### 4. 更新博客文章

- **URL**: `/api/blogs/:id`
- **Method**: `PUT`
- **Description**: 更新指定ID的博客文章
- **URL Parameters**:
  - `id`: 博客文章ID
- **Request Body**:
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "title": "更新后的博客标题",
      "content": "更新后的博客内容（Markdown格式）",
      "author": "作者名称"
    }
    ```
- **Success Response**:
  - **Code**: 200
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "data": {
        "id": 1,
        "title": "更新后的博客标题",
        "content": "更新后的博客内容（Markdown格式）",
        "author": "作者名称",
        "createdAt": "2024-01-15T10:00:00Z",
        "updatedAt": "2024-01-20T15:00:00Z"
      }
    }
    ```
- **Error Response**:
  - **Code**: 400
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "标题和内容不能为空"
    }
    ```
  - **Code**: 404
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "博客不存在"
    }
    ```
  - **Code**: 500
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "更新博客失败"
    }
    ```

#### 5. 删除博客文章

- **URL**: `/api/blogs/:id`
- **Method**: `DELETE`
- **Description**: 删除指定ID的博客文章
- **URL Parameters**:
  - `id`: 博客文章ID
- **Success Response**:
  - **Code**: 200
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "data": {
        "message": "删除成功"
      }
    }
    ```
- **Error Response**:
  - **Code**: 404
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "博客不存在"
    }
    ```
  - **Code**: 500
  - **Content-Type**: application/json
  - **Body**:
    ```json
    {
      "error": "删除博客失败"
    }
    ```

## 数据模型

### 博客文章 (Blog)

```javascript
{
  id: Number,         // 唯一标识符
  title: String,      // 标题
  content: String,    // 内容（Markdown格式）
  author: String,     // 作者
  createdAt: String,  // 创建时间（ISO 8601格式）
  updatedAt: String   // 更新时间（ISO 8601格式）
}
```

## 部署建议

1. 后端API服务可以使用Node.js + Express/Koa、Python + Flask/Django、Go等任何语言和框架实现
2. 数据库可以使用MySQL、PostgreSQL、MongoDB等关系型或非关系型数据库
3. 建议在生产环境中添加HTTPS支持和适当的认证机制
4. 考虑添加数据缓存层以提高性能
5. 实现适当的日志记录和错误处理机制

## 注意事项

1. 本API文档基于前端需求制定，后端实现时可根据实际情况进行调整
2. 所有日期时间格式应使用ISO 8601标准
3. 建议添加适当的API版本控制机制
4. 为提高安全性，建议实现输入验证和防XSS攻击