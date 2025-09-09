<template>
  <div class="blog-detail">
    <div v-if="loading" class="loading">加载中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else-if="blog" class="blog-content">
      <div class="actions">
        <router-link to="/" class="back-button">返回列表</router-link>
        <router-link :to="{ name: 'BlogEdit', params: { id: blog.id } }" class="edit-button">编辑</router-link>
      </div>
      
      <h1>{{ blog.title }}</h1>
      
      <div class="meta">
        <span>作者: {{ blog.author || '匿名' }}</span>
        <span class="separator">|</span>
        <span>发布时间: {{ formatDate(blog.createdAt) }}</span>
        <span v-if="blog.updatedAt !== blog.createdAt" class="separator">|</span>
        <span v-if="blog.updatedAt !== blog.createdAt">更新时间: {{ formatDate(blog.updatedAt) }}</span>
      </div>
      
      <div class="content" v-html="renderedContent"></div>
    </div>
    
    <div v-else class="not-found">博客不存在</div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import MarkdownIt from 'markdown-it'
import blogApi from '../services/api'

const route = useRoute()
const blog = ref(null)
const loading = ref(false)
const error = ref('')

// 创建Markdown解析器实例
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true
})

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// 计算渲染后的Markdown内容
const renderedContent = computed(() => {
  if (!blog.value?.content) return ''
  return md.render(blog.value.content)
})

// 获取博客详情
const fetchBlogDetail = async () => {
  const id = route.params.id
  loading.value = true
  error.value = ''
  try {
    // 调用真实API获取博客详情
    const response = await blogApi.getBlogById(id)
    blog.value = response.data
    
    if (!blog.value) {
      error.value = '博客不存在'
    }
  } catch (err) {
    error.value = '获取博客详情失败'
    console.error('获取博客详情失败:', err)
  } finally {
    loading.value = false
  }
}

// 组件挂载时获取博客详情
onMounted(() => {
  fetchBlogDetail()
})

// 监听路由参数变化，重新获取数据
route.params.id && fetchBlogDetail()
</script>

<style scoped>
.blog-detail {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.loading, .not-found, .error {
  padding: 20px;
  text-align: center;
  color: #666;
}

.error {
  color: #f44336;
}

.actions {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.back-button {
  background-color: #666;
  color: white;
  padding: 8px 15px;
  text-decoration: none;
  border-radius: 4px;
  font-size: 0.9em;
  transition: background-color 0.3s;
}

.back-button:hover {
  background-color: #555;
}

.edit-button {
  background-color: #2196F3;
  color: white;
  padding: 8px 15px;
  text-decoration: none;
  border-radius: 4px;
  font-size: 0.9em;
  transition: background-color 0.3s;
}

.edit-button:hover {
  background-color: #0b7dda;
}

.blog-content h1 {
  color: #333;
  margin-top: 0;
  margin-bottom: 15px;
}

.meta {
  color: #666;
  font-size: 0.9em;
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.separator {
  margin: 0 10px;
}

.content {
  color: #333;
  line-height: 1.8;
}

/* Markdown内容样式 */
.content h1,
.content h2,
.content h3,
.content h4,
.content h5,
.content h6 {
  color: #2c3e50;
  margin-top: 1.5em;
  margin-bottom: 0.5em;
}

.content p {
  margin-bottom: 1em;
}

.content a {
  color: #3498db;
  text-decoration: none;
}

.content a:hover {
  text-decoration: underline;
}

.content ul,
.content ol {
  margin-bottom: 1em;
  padding-left: 1.5em;
}

.content li {
  margin-bottom: 0.5em;
}

.content code {
  background-color: #f5f5f5;
  padding: 2px 4px;
  border-radius: 3px;
  font-family: 'Courier New', Courier, monospace;
}

.content pre {
  background-color: #f5f5f5;
  padding: 15px;
  border-radius: 4px;
  overflow-x: auto;
  margin-bottom: 1em;
}

.content pre code {
  background-color: transparent;
  padding: 0;
}

.content blockquote {
  border-left: 4px solid #3498db;
  padding-left: 15px;
  margin-left: 0;
  color: #666;
}
</style>