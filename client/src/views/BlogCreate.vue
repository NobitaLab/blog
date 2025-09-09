<template>
  <div class="blog-create">
    <h1>创建新博客</h1>
    
    <div v-if="loading" class="loading">处理中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else class="create-form">
      <div class="form-group">
        <label for="title">标题</label>
        <input
          id="title"
          v-model="formData.title"
          type="text"
          placeholder="请输入博客标题"
          class="form-input"
        />
      </div>
      
      <div class="form-group">
        <label for="author">作者</label>
        <input
          id="author"
          v-model="formData.author"
          type="text"
          placeholder="请输入作者名称"
          class="form-input"
        />
      </div>
      
      <div class="editor-toolbar">
        <button
          @click="switchEditor('markdown')"
          :class="['editor-btn', { active: editorMode === 'markdown' }]"
        >
          Markdown
        </button>
        <button
          @click="switchEditor('preview')"
          :class="['editor-btn', { active: editorMode === 'preview' }]"
        >
          预览
        </button>
        <button
          @click="switchEditor('split')"
          :class="['editor-btn', { active: editorMode === 'split' }]"
        >
          分屏
        </button>
      </div>
      
      <div class="editor-container">
        <div v-if="editorMode === 'markdown' || editorMode === 'split'" class="markdown-editor">
          <textarea
            v-model="formData.content"
            placeholder="请输入博客内容（支持Markdown格式）"
            class="markdown-textarea"
            rows="20"
          ></textarea>
        </div>
        
        <div v-if="editorMode === 'preview' || editorMode === 'split'" class="preview-container">
          <div class="preview-content" v-html="renderedContent"></div>
        </div>
      </div>
      
      <div class="form-actions">
        <button @click="createBlog" class="create-button">创建</button>
        <router-link to="/" class="cancel-button">取消</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import MarkdownIt from 'markdown-it'
import blogApi from '../services/api'

const router = useRouter()
const loading = ref(false)
const error = ref('')
const editorMode = ref('split') // markdown, preview, split

// 表单数据
const formData = ref({
  title: '',
  author: '',
  content: ''
})

// 创建Markdown解析器实例
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true
})

// 计算渲染后的Markdown内容
const renderedContent = computed(() => {
  if (!formData.value.content) return '<p>预览内容将显示在这里</p>'
  return md.render(formData.value.content)
})

// 切换编辑器模式
const switchEditor = (mode) => {
  editorMode.value = mode
}

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
      author: formData.value.author.trim() || '匿名',
      content: formData.value.content.trim()
    }
    
    // 调用真实API创建博客
    const response = await blogApi.createBlog(data)
    alert('博客创建成功')
    router.push(`/blog/${response.data.id}`)
  } catch (err) {
    error.value = '创建博客失败'
    alert(error.value)
    console.error(error.value, err)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.blog-create {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  margin-bottom: 20px;
}

.loading, .error {
  padding: 20px;
  text-align: center;
  color: #666;
}

.error {
  color: #f44336;
}

.create-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-weight: bold;
  color: #555;
}

.form-input {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.form-input:focus {
  outline: none;
  border-color: #2196F3;
  box-shadow: 0 0 0 2px rgba(33, 150, 243, 0.2);
}

.editor-toolbar {
  display: flex;
  gap: 10px;
  background-color: #f5f5f5;
  padding: 10px;
  border-radius: 4px 4px 0 0;
  border-bottom: 1px solid #ddd;
}

.editor-btn {
  padding: 8px 15px;
  border: 1px solid #ddd;
  background-color: white;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s;
}

.editor-btn:hover {
  background-color: #f0f0f0;
}

.editor-btn.active {
  background-color: #2196F3;
  color: white;
  border-color: #2196F3;
}

.editor-container {
  display: flex;
  border: 1px solid #ddd;
  border-top: none;
  border-radius: 0 0 4px 4px;
  overflow: hidden;
}

.markdown-editor,
.preview-container {
  flex: 1;
  min-height: 500px;
}

.markdown-editor {
  border-right: 1px solid #ddd;
}

.markdown-textarea {
  width: 100%;
  height: 100%;
  padding: 15px;
  border: none;
  resize: none;
  font-family: 'Courier New', Courier, monospace;
  font-size: 14px;
  line-height: 1.6;
}

.markdown-textarea:focus {
  outline: none;
}

.preview-container {
  padding: 15px;
  overflow-y: auto;
}

.preview-content {
  color: #333;
  line-height: 1.8;
}

/* Markdown内容样式 */
.preview-content h1,
.preview-content h2,
.preview-content h3,
.preview-content h4,
.preview-content h5,
.preview-content h6 {
  color: #2c3e50;
  margin-top: 1.5em;
  margin-bottom: 0.5em;
}

.preview-content p {
  margin-bottom: 1em;
}

.preview-content a {
  color: #3498db;
  text-decoration: none;
}

.preview-content a:hover {
  text-decoration: underline;
}

.preview-content ul,
.preview-content ol {
  margin-bottom: 1em;
  padding-left: 1.5em;
}

.preview-content li {
  margin-bottom: 0.5em;
}

.preview-content code {
  background-color: #f5f5f5;
  padding: 2px 4px;
  border-radius: 3px;
  font-family: 'Courier New', Courier, monospace;
}

.preview-content pre {
  background-color: #f5f5f5;
  padding: 15px;
  border-radius: 4px;
  overflow-x: auto;
  margin-bottom: 1em;
}

.preview-content pre code {
  background-color: transparent;
  padding: 0;
}

.preview-content blockquote {
  border-left: 4px solid #3498db;
  padding-left: 15px;
  margin-left: 0;
  color: #666;
}

.form-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.create-button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s;
}

.create-button:hover {
  background-color: #45a049;
}

.cancel-button {
  background-color: #666;
  color: white;
  padding: 10px 20px;
  text-decoration: none;
  border-radius: 4px;
  font-weight: bold;
  transition: background-color 0.3s;
}

.cancel-button:hover {
  background-color: #555;
}
</style>