import axios from 'axios'

// 创建axios实例
const api = axios.create({
  baseURL: '/api', // 后端API的基础URL
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 这里可以添加token等认证信息
    // const token = localStorage.getItem('token')
    // if (token) {
    //   config.headers.Authorization = `Bearer ${token}`
    // }
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
  deleteBlog: (id) => api.delete(`/blogs/${id}`)
}

export default blogApi