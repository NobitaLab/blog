<template>
  <div class="blog-list">
    <h1>我的博客</h1>
    <div class="actions">
      <router-link to="/create" class="create-button">创建新博客</router-link>
    </div>
    
    <div v-if="loading" class="loading">加载中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else-if="blogs.length === 0" class="empty">暂无博客文章</div>
    
    <div v-else class="blogs-container">
      <div v-for="blog in blogs" :key="blog.id" class="blog-item">
        <h2><router-link :to="{ name: 'BlogDetail', params: { id: blog.id } }">{{ blog.title }}</router-link></h2>
        <p class="meta">
          <span>{{ formatDate(blog.createdAt) }}</span>
          <span class="separator">|</span>
          <span>{{ blog.author || '匿名' }}</span>
        </p>
        <p class="excerpt">{{ blog.content.substring(0, 150) }}...</p>
        <div class="blog-actions">
          <router-link :to="{ name: 'BlogEdit', params: { id: blog.id } }" class="edit-button">编辑</router-link>
          <button @click="deleteBlog(blog.id)" class="delete-button">删除</button>
          <!-- 解释：
            - @click是vue的事件绑定指令（简写形式，等价于 v-on:click），用于给按钮绑定 “点击事件”。当按钮被点击时，会执行 deleteBlog(blog.id) 函数
            - deleteBlog(blog.id) 是一个自定义的方法，用于删除指定id的博客文章
            - 这里的 blog.id 是当前循环遍历到的博客文章的id，用于标识要删除的具体文章
          -->
        </div>
      </div>
    </div>
    
    <!-- 分页控件 -->
    <div class="pagination" v-if="total > pageSize">
      <button 
        @click="currentPage--; fetchBlogs()"
        :disabled="currentPage <= 1"
        class="page-button"
      >
        上一页
      </button>
      <span class="page-info">
        第 {{ currentPage }} 页 / 共 {{ Math.ceil(total / pageSize) }} 页
      </span>
      <button 
        @click="currentPage++; fetchBlogs()"
        :disabled="currentPage >= Math.ceil(total / pageSize)"
        class="page-button"
      >
        下一页
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { router } from '../main'
import blogApi from '../services/api'

const blogs = ref([])
const loading = ref(false)
const error = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// 获取所有博客文章
const fetchBlogs = async () => {
  loading.value = true
  error.value = ''
  try {
    // 调用真实API获取博客列表
    const response = await blogApi.getAllBlogs({ page: currentPage.value, limit: pageSize.value })
    blogs.value = response.data
    total.value = response.total
    currentPage.value = response.page
    pageSize.value = response.limit
  } catch (err) {
    error.value = '获取博客列表失败'
    console.error('获取博客列表失败:', err)
  } finally {
    loading.value = false
  }
}

// 删除博客文章
const deleteBlog = async (id) => {
  if (confirm('确定要删除这篇博客吗？')) {
    try {
      // 调用真实API删除博客
      await blogApi.deleteBlog(id)
      
      // 删除成功后刷新博客列表
      await fetchBlogs()
      
      // 提示删除成功
      alert('博客删除成功')
    } catch (err) {
      alert('删除失败，请重试')
      console.error('删除博客失败:', err)
    }
  }
}

// 组件挂载时获取博客列表
onMounted(fetchBlogs)
</script>

<style scoped>
.blog-list {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  margin-bottom: 20px;
}

.actions {
  margin-bottom: 20px;
}

.create-button {
  display: inline-block;
  background-color: #4CAF50;
  color: white;
  padding: 10px 15px;
  text-decoration: none;
  border-radius: 4px;
  font-weight: bold;
  transition: background-color 0.3s;
}

.create-button:hover {
  background-color: #45a049;
}

.loading, .empty, .error {
  padding: 20px;
  text-align: center;
  color: #666;
}

.error {
  color: #f44336;
}

.blogs-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.blog-item {
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}

.blog-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
}

.blog-item h2 {
  margin-top: 0;
  color: #333;
}

.blog-item h2 a {
  color: #333;
  text-decoration: none;
  transition: color 0.3s;
}

.blog-item h2 a:hover {
  color: #0066cc;
}

.meta {
  color: #666;
  font-size: 0.9em;
  margin-bottom: 10px;
}

.separator {
  margin: 0 10px;
}

.excerpt {
  color: #444;
  line-height: 1.6;
  margin-bottom: 15px;
}

.blog-actions {
  display: flex;
  gap: 10px;
}

.edit-button {
  background-color: #2196F3;
  color: white;
  padding: 6px 12px;
  text-decoration: none;
  border-radius: 4px;
  font-size: 0.9em;
  transition: background-color 0.3s;
}

.edit-button:hover {
  background-color: #0b7dda;
}

.delete-button {
  background-color: #f44336;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9em;
  transition: background-color 0.3s;
}

.delete-button:hover {
  background-color: #da190b;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 15px;
}

.page-button {
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

.page-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.page-info {
  color: #666;
}
</style>