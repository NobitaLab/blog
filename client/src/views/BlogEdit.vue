<template>
  <div class="blog-edit">
    
    <div v-if="loading" class="loading">加载中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else class="edit-form">
      <!-- 使用高级编辑器组件 -->
      <AdvancedEditor
        v-model:title="formData.title"
        v-model="formData.content"
        :disabled="isEditMode && !canEdit"
        @save="saveBlog"
      />
    </div>
    
    <!-- 权限提示 -->
    <div v-if="isEditMode && !canEdit && !loading" class="permission-tip">
      您只能查看这篇博客，没有编辑权限。<br>
      只有博客作者或管理员才能编辑此博客。
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AdvancedEditor from '../components/AdvancedEditor.vue'
import blogApi from '../services/api.js'

const route = useRoute()
const router = useRouter()
const isEditMode = ref(false)
const loading = ref(false)
const error = ref('')
const currentUser = ref(null)
const isAdmin = ref(false)
const canEdit = ref(true)
const blogData = ref(null)

// 表单数据
const formData = ref({
  title: '',
  content: ''
})

// 检查用户认证状态
const checkUserStatus = async () => {
  try {
    const response = await blogApi.getCurrentUser()
    if (response.data) {
      currentUser.value = response.data
      isAdmin.value = response.data.role === 'admin'
      return true
    }
  } catch (err) {
    console.error('获取用户信息失败:', err)
  }
  currentUser.value = null
  isAdmin.value = false
  return false
}

// 检查编辑权限
const checkEditPermission = () => {
  if (!blogData.value || !currentUser.value) {
    return false
  }
  // 管理员可以编辑所有博客
  if (isAdmin.value) {
    return true
  }
  // 普通用户只能编辑自己的博客
  return blogData.value.author === currentUser.value.username
}

// 获取博客详情（编辑模式）
const fetchBlogDetail = async () => {
  const id = route.params.id
  if (id) {
    isEditMode.value = true
    loading.value = true
    error.value = ''
    try {
      // 检查用户认证状态
      const isAuthenticated = await checkUserStatus()
      if (!isAuthenticated) {
        error.value = '请先登录'
        loading.value = false
        return
      }
      
      // 调用真实API获取博客详情
      const response = await blogApi.getBlogById(id)
      blogData.value = response.data
      formData.value = {
        title: response.data.title,
        content: response.data.content
      }
      
      // 检查编辑权限
      canEdit.value = checkEditPermission()
      if (!canEdit.value) {
        error.value = '您没有权限编辑这篇博客'
        return
      }
    } catch (err) {
      error.value = '获取博客详情失败'
      console.error('获取博客详情失败:', err)
    } finally {
      loading.value = false
    }
  } else {
    // 创建模式，检查用户认证状态
    checkUserStatus()
  }
}

// 保存博客
const saveBlog = async () => {
  // 表单验证
  if (!formData.value.title.trim()) {
    alert('请输入标题')
    return
  }
  if (!formData.value.content.trim()) {
    alert('请输入内容')
    return
  }
  
  // 检查是否有权限保存
  if (isEditMode.value && !canEdit.value) {
    alert('您没有权限保存这篇博客')
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const data = {
      title: formData.value.title.trim(),
      content: formData.value.content.trim()
    }
    
    if (isEditMode.value) {
      // 调用真实API更新博客
      await blogApi.updateBlog(route.params.id, data)
      alert('博客更新成功')
      router.push(`/blog/${route.params.id}`)
    } else {
      // 调用真实API创建博客
      const response = await blogApi.createBlog(data)
      alert('博客创建成功')
      router.push(`/blog/${response.data.id}`)
    }
  } catch (err) {
    error.value = err.response?.data?.error || (isEditMode.value ? '更新博客失败' : '创建博客失败')
    alert(error.value)
    console.error(isEditMode.value ? '更新博客失败:' : '创建博客失败:', err)
  } finally {
    loading.value = false
  }
}

// 组件挂载时，如果是编辑模式则获取博客详情
onMounted(() => {
  fetchBlogDetail()
})
</script>

<style scoped>
.blog-edit {
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

.edit-form {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0;
  width: 100%;
  overflow: hidden;
  min-height: 0;
}

.permission-tip {
  margin-top: 20px;
  padding: 15px;
  background-color: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: 4px;
  color: #856404;
  text-align: center;
  line-height: 1.6;
}
</style>