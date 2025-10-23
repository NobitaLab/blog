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
      
      <div class="markdown-body" v-html="renderedContent"></div>
    </div>
    
    <div v-else class="not-found">博客不存在</div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js'
import mermaid from 'mermaid'
import taskLists from 'markdown-it-task-lists'
import ins from 'markdown-it-ins'
import footnote from 'markdown-it-footnote'
import abbr from 'markdown-it-abbr'
import mark from 'markdown-it-mark'
import sub from 'markdown-it-sub'
import sup from 'markdown-it-sup'
import deflist from 'markdown-it-deflist'
import blogApi from '../services/api'

const route = useRoute()
const blog = ref(null)
const loading = ref(false)
const error = ref('')

// 创建MarkdownIt实例并配置插件（与AdvancedEditor.vue保持一致）
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
  breaks: true, // 支持换行符转换为<br>
  highlight: function(code, lang) {
    console.log('highlight函数被调用:', { code: code.substring(0, 50) + '...', lang });
    
    if (lang && hljs.getLanguage(lang)) {
      try {
        const result = hljs.highlight(code, { language: lang }).value
        console.log('highlight结果:', result.substring(0, 100) + '...');
        return result
      } catch (err) {
        console.log('highlight失败，返回原代码');
        return code
      }
    }
    try {
      const result = hljs.highlightAuto(code).value
      console.log('auto highlight结果:', result.substring(0, 100) + '...');
      return result
    } catch (err) {
      console.log('auto highlight失败，返回原代码');
      return code
    }
  }
})

// 测试markdown渲染
console.log('测试markdown渲染:');
const testContent = '```javascript\nconsole.log("test");\n```';
const testResult = md.render(testContent);
console.log('测试结果:', testResult);

// 使用插件扩展功能 - 不同内容类型分配给相应插件处理
// 为每个插件添加独立的错误处理，确保一个插件失败不影响其他插件

// 任务列表内容处理
try {
  md.use(taskLists, { 
    enabled: true, // 启用任务列表
    label: true,  // 允许点击标签切换状态
    clickable: true // 允许点击复选框切换状态
  })
} catch (e) {
  // taskLists插件加载失败，跳过
}

// 文本格式化内容处理
try {
  md.use(ins)     // 处理插入线内容
} catch (e) {
  // ins插件加载失败，跳过
}

try {
  md.use(mark)    // 处理标记内容
} catch (e) {
  // mark插件加载失败，跳过
}

try {
  md.use(sub)     // 处理下标内容
} catch (e) {
  // sub插件加载失败，跳过
}

try {
  md.use(sup)     // 处理上标内容
} catch (e) {
  // sup插件加载失败，跳过
}

// 结构化内容处理
try {
  md.use(abbr)    // 处理缩写定义内容
} catch (e) {
  // abbr插件加载失败，跳过
}

try {
  md.use(deflist) // 处理定义列表内容
} catch (e) {
  // deflist插件加载失败，跳过
}

try {
  md.use(footnote) // 处理脚注引用和定义内容
} catch (e) {
  // footnote插件加载失败，跳过
}

// 为mermaid图表添加自定义规则处理
md.renderer.rules.fence = function(tokens, idx, options, env, self) {
  const token = tokens[idx];
  const info = token.info ? token.info.trim() : '';
  
  console.log('处理fence规则:', { info, content: token.content.substring(0, 50) + '...' });
  
  if (info.startsWith('mermaid')) {
    // 对于mermaid图表，直接创建带有mermaid类的div，让mermaid库自动识别
    const code = token.content.trim();
    const id = `mermaid-${Date.now()}-${Math.floor(Math.random() * 1000)}`;
    return `<div class="mermaid-chart-container" style="margin: 16px 0; padding: 16px; background-color: #f8f9fa; border-radius: 4px;">
      <div class="mermaid" id="${id}">${code}</div>
    </div>`;
  }
  
  // 对于其他代码块，使用默认渲染（包含代码高亮）
  const result = self.renderToken(tokens, idx, options);
  console.log('fence渲染结果:', result.substring(0, 100) + '...');
  return result;
}

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
  
  try {
    // 添加调试信息
    console.log('开始渲染markdown内容:', blog.value.content.substring(0, 100) + '...')
    
    // 测试简单的markdown渲染
    const testContent = '```javascript\nconsole.log("test");\n```'
    const testHtml = md.render(testContent)
    console.log('测试代码块渲染:', testHtml)
    
    const html = md.render(blog.value.content)
    
    console.log('渲染后的HTML:', html.substring(0, 200) + '...')
    
    // 在下一个tick中初始化mermaid图表
    nextTick(() => {
      initMermaidCharts()
    })
    
    return html
  } catch (error) {
    console.error('Markdown渲染错误:', error)
    return blog.value.content
  }
})

// 初始化mermaid图表
const initMermaidCharts = () => {
  const mermaidElements = document.querySelectorAll('.mermaid')
  if (mermaidElements.length > 0) {
    mermaid.init(undefined, mermaidElements)
  }
}

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
  background-color: #f5f5f5;
  min-height: calc(100vh - 90px); /* 减去header高度和padding */
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

/* 使用github-markdown-css提供的样式，只需要添加一些自定义样式 */
.markdown-body {
  max-width: 800px;
  margin: 0 auto;
  padding: 24px;
  box-sizing: border-box;
  background-color: #f5f5f5; /* 与页面背景色保持一致 */
}

/* 任务列表样式 */
.markdown-body .task-list-item {
  list-style-type: none;
}

.markdown-body .task-list-item input[type="checkbox"] {
  margin-right: 8px;
}

/* Mermaid图表样式 */
.markdown-body .mermaid {
  margin: 24px 0;
  overflow-x: auto;
  text-align: center;
}

.markdown-body .mermaid-chart-container {
  margin: 16px 0;
  padding: 16px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.markdown-body .mermaid-rendered {
  text-align: center;
  margin: 24px 0;
  padding: 16px;
  background: #f6f8fa;
  border-radius: 6px;
}


</style>