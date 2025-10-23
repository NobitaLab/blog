import axios from 'axios'
import { router } from '../main.js'

// 创建axios实例
const api = axios.create({
  baseURL: 'http://localhost:8080/api', // 后端API的基础地址
  timeout: 5000, // 降低超时时间从10秒到5秒，提高用户体验
  headers: {
    'Content-Type': 'application/json'
  },
  // 允许跨域请求携带凭证信息
  withCredentials: true 
})

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 添加token等认证信息
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    console.error('API请求错误:', error)
    
    // 统一处理认证失败
    if (error.response && error.response.status === 401) {
      // 清除本地存储的token和用户信息
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      // 跳转到登录页面
      router.push('/login')
    }
    
    return Promise.reject(error)
  }
)

// 博客文章相关API
const blogApi = {
  // 获取所有博客文章
  getAllBlogs: (params) => api.get('/blogs', { params }),
  
  // 获取单篇博客文章
  getBlogById: (id) => api.get(`/blogs/${id}`),
  
  // 创建新的博客文章
  createBlog: (data) => api.post('/blogs', data),
  
  // 更新博客文章
  updateBlog: (id, data) => api.put(`/blogs/${id}`, data),
  
  // 删除博客文章
  deleteBlog: (id) => api.delete(`/blogs/${id}`),
  
  // 用户认证相关API
  // 用户注册
  register: (data) => api.post('/users/register', data),
  
  // 用户登录
  login: (data) => api.post('/users/login', data),
  
  // 用户注销
  logout: () => api.post('/users/logout'),
  
  // 获取当前用户信息
  getCurrentUser: () => api.get('/users/me'),
  
  // 管理员API
  // 获取所有用户信息
  getAllUsers: () => api.get('/admin/users'),
  
  // 更新用户信息
  updateUser: (id, data) => api.put(`/admin/users/${id}`, data),
  
  // 图片上传API
  uploadImage: (formData) => {
    return api.post('/upload/image', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  },
  
  // 图片删除API
  deleteImage: (filename) => {
    return api.post('/upload/image/delete', new URLSearchParams({
      filename: filename
    }), {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    })
  }
}

export default blogApi

export { router }