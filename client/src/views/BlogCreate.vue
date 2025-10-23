<template>
  <div class="blog-create">
    
    <div v-if="loading" class="loading">处理中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else class="create-form">
      <!-- 使用高级编辑器组件 -->
      <AdvancedEditor
        v-model:title="formData.title"
        v-model="formData.content"
        :disabled="false"
        @save="createBlog"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AdvancedEditor from '../components/AdvancedEditor.vue'
import blogApi from '../services/api.js'

const router = useRouter()
const loading = ref(false)
const error = ref('')

// 表单数据
const formData = ref({
  title: '',
  content: ''
})

// 创建博客
const createBlog = async () => {
  // 表单验证
  if (!formData.value.title.trim()) {
    alert('请输入标题')
    return
  }
  if (!formData.value.content.trim()) {
    alert('请输入内容')
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const data = {
      title: formData.value.title.trim(),
      content: formData.value.content.trim()
    }
    
    // 调用真实API创建博客
    const response = await blogApi.createBlog(data)
    alert('博客创建成功')
    router.push(`/blog/${response.data.id}`)
  } catch (err) {
    error.value = err.response?.data?.error || '创建博客失败'
    alert(error.value)
    console.error('创建博客失败:', err)
  } finally {
    loading.value = false
  }
}

// 组件挂载时检查用户状态
onMounted(() => {
  // 检查用户是否已登录
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
    return
  }
})
</script>

<style scoped>
.blog-create {
  width: 100%;
  height: calc(100vh - 50px);
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  position: fixed;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
}

.loading, .error {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  text-align: center;
  color: #666;
}

.error {
  color: #f44336;
}

.create-form {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0;
  width: 100%;
  overflow: hidden;
  min-height: 0;
}
</style>