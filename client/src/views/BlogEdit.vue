<template>
  <div class="blog-edit">
    <h1>{{ isEditMode ? '编辑博客' : '创建博客' }}</h1>
    
    <div v-if="loading" class="loading">加载中...</div>
    
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <div v-else class="edit-form">
      <div class="form-group">
        <label for="title">标题</label>
        <!-- 解释：
          - <label for="title"> 是标签，用于描述输入框的作用
          - for="title" 与输入框的 id="title" 关联，点击标签时会自动聚焦到输入框
        -->
        <input
          id="title"
          v-model="formData.title"
          type="text"
          placeholder="请输入博客标题"
          class="form-input"
        />
        <!-- 解释：
            - v-model 指令将输入框的值与 formData.title 双向绑定
            - 当输入框的值改变时，formData.title 也会更新
            - 当 formData.title 改变时，输入框的值也会更新
          -->
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
          <!-- 解释：
            - @click是vue的事件绑定指令（简写形式，等价于 v-on:click），用于给按钮绑定 “点击事件”。当按钮被点击时，会执行 switchEditor('markdown') 函数
            - :class 是 Vue 的动态类绑定指令（简写形式，等价于 v-bind:class），用于根据数据状态动态给元素添加 / 移除 CSS 类。
            - 这里使用了数组语法（数组中可包含固定类名和动态类对象），固定类名 'editor-btn' 会始终添加，
            - 而 { active: editorMode === 'markdown' } 会根据 editorMode 的值动态添加或移除 active 类名
            - 当 editorMode 等于 'markdown' 时，添加 active 类名，否则移除 active 类名
          -->
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
          <!-- 解释：
            - textarea标签：用于多行文本输入
            - v-model指令：将输入框的值与 formData.content 双向绑定
            - placeholder属性：指定输入框为空时显示的提示文本
            - class属性：添加 markdown-textarea 类，用于样式化
            - rows属性：指定文本区域的行数
          -->
        </div>
        
        <div v-if="editorMode === 'preview' || editorMode === 'split'" class="preview-container">
          <div class="preview-content" v-html="renderedContent"></div>
          <!-- 解释：
            - v-html指令：将渲染后的Markdown内容插入到div中，v-html 是 Vue 提供的用于动态渲染 HTML 内容的指令，
            - 它的核心作用是：将绑定的数据（字符串）解析为 HTML 标签并渲染到页面中，而不是像 {{ }} 插值语法那样将内容作为纯文本显示。
          -->
        </div>
      </div>
      
      <div class="form-actions">
        <button @click="saveBlog" class="save-button">保存</button>
        <router-link to="/" class="cancel-button">取消</router-link>
        <!-- 解释：
          - <router-link> 和 <button> 适用场景及渲染场景差异：
            - <router-link> 是 Vue 提供的组件，用于路由跳转，适用于页面跳转,无额外逻辑场景，渲染为 <a> 标签
            - <button> 是 HTML 元素，用于触发事件，适用于提交表单、执行操作等需要执行复杂逻辑的操作场景，渲染为 <button> 标签
        -->
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import MarkdownIt from 'markdown-it'
import blogApi from '../services/api'

const route = useRoute()
const router = useRouter()
const isEditMode = ref(false)
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

// 获取博客详情（编辑模式）
const fetchBlogDetail = async () => {
  const id = route.params.id
  if (id) {
    isEditMode.value = true
    loading.value = true
    error.value = ''
    try {
      // 调用真实API获取博客详情
      const response = await blogApi.getBlogById(id)
      formData.value = {
        title: response.data.title,
        author: response.data.author,
        content: response.data.content
      }
      formData.value = {
          title: response.data.title,
          author: response.data.author,
          content: response.data.content
        }
    } catch (err) {
      error.value = '获取博客详情失败'
      console.error('获取博客详情失败:', err)
    } finally {
      loading.value = false
    }
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
  
  loading.value = true
  error.value = ''
  try {
    const data = {
      title: formData.value.title.trim(),
      author: formData.value.author.trim() || '匿名',
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
      router.push('/')
    }
  } catch (err) {
    error.value = isEditMode.value ? '更新博客失败' : '创建博客失败'
    alert(error.value)
    console.error(error.value, err)
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

.edit-form {
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

.save-button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s;
}

.save-button:hover {
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