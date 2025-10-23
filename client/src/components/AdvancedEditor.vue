<template>
  <div class="juejin-editor">
    <!-- 顶部标题栏 -->
    <header class="juejin-header">
      <div class="header-left">
        <input 
          v-model="title" 
          placeholder="输入文章标题..." 
          class="title-input"
          maxlength="80"
          spellcheck="false"
          @input="handleTitleChange"
        />
      </div>
      <div class="header-right">
        <div class="status-text">已保存</div>
        <button class="draft-btn" @click="handleDraft">草稿箱</button>
        <div class="publish-popup">
          <button class="publish-btn" @click="handlePublish">发布</button>
        </div>
        <div class="editor-switcher">
          <div class="toggle-btn">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path d="M4.7998 9.6001L19.1998 9.6001L15.1998 5.6001" stroke="#86909C" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M19.1997 14.3999L4.79971 14.3999L8.79971 18.3999" stroke="#86909C" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </div>
        </div>
        <nav class="user-navigator">
          <div class="user-avatar"></div>
        </nav>
      </div>
    </header>

    <!-- 主编辑器区域 -->
    <div class="juejin-main">
      <div class="bytemd bytemd-split">
        <!-- 工具栏 -->
        <div class="bytemd-toolbar">
          <div class="bytemd-toolbar-left">
            <!-- 标题 -->
            <div class="bytemd-toolbar-item relative">
              <div class="bytemd-toolbar-icon" title="标题" @mouseenter="showHeadingDropdown = true">
                <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 5v38M36 5v38M12 24h24"/>
                </svg>
              </div>
              <!-- 标题级别下拉菜单 -->
              <div v-if="showHeadingDropdown" class="heading-dropdown" 
                   @mouseenter="showHeadingDropdown = true"
                   @mouseleave="showHeadingDropdown = false" 
                   @mouseover.stop>
                <div v-for="level in 6" :key="level" 
                     class="heading-option" 
                     @click.stop="insertHeadingLevel(level)">
                  <span class="heading-level">H{{ level }}</span>
                  <span class="heading-preview">标题{{ level }}</span>
                </div>
              </div>
            </div>
            <!-- 粗体 -->
            <div class="bytemd-toolbar-icon" title="粗体" @click="insertBold">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24c5.506 0 9.969-4.477 9.969-10S29.506 4 24 4H11v20h13ZM28.031 44C33.537 44 38 39.523 38 34s-4.463-10-9.969-10H11v20h17.031Z" clip-rule="evenodd"/>
              </svg>
            </div>
            <!-- 斜体 -->
            <div class="bytemd-toolbar-icon" title="斜体" @click="insertItalic">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 6h16M12 42h16M29 5.952 19 42"/>
              </svg>
            </div>
            <!-- 引用 -->
            <div class="bytemd-toolbar-icon" title="引用" @click="insertQuote">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path fill="currentColor" fill-rule="evenodd" d="M18.853 9.116C11.323 13.952 7.14 19.58 6.303 26.003 5 36 13.94 40.893 18.47 36.497 23 32.1 20.285 26.52 17.005 24.994c-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.016-7.69 9.116-10.322a.749.749 0 0 0 .114-1.02L20.285 9.3c-.44-.572-.862-.55-1.432-.185ZM38.679 9.116c-7.53 4.836-11.714 10.465-12.55 16.887-1.303 9.997 7.637 14.89 12.167 10.494 4.53-4.397 1.815-9.977-1.466-11.503-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.017-7.69 9.117-10.322a.749.749 0 0 0 .113-1.02L40.11 9.3c-.44-.572-.862-.55-1.431-.185Z" clip-rule="evenodd"/>
              </svg>
            </div>
            <!-- 链接 -->
            <div class="bytemd-toolbar-icon" title="链接" @click="insertLink">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26.24 16.373-9.14-9.14c-2.661-2.661-7.035-2.603-9.768.131-2.734 2.734-2.793 7.107-.131 9.768l7.935 7.936M32.903 23.003l7.935 7.935c2.661 2.662 2.603 7.035-.13 9.769-2.735 2.734-7.108 2.792-9.77.13l-9.14-9.14"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26.11 26.142c2.733-2.734 2.792-7.108.13-9.769M21.799 21.798c-2.734 2.734-2.792 7.108-.131 9.769"/>
              </svg>
            </div>
            <!-- 图片 -->
            <div class="bytemd-toolbar-icon" title="图片" @click="insertImage">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 10a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V10Z" clip-rule="evenodd"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14.5 18a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>
                <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m15 24 5 4 6-7 17 13v4a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-4l10-10Z"/>
              </svg>
            </div>
            <!-- 代码 -->
            <div class="bytemd-toolbar-icon" title="代码" @click="insertInlineCode">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 13 4 25.432 16 37M32 13l12 12.432L32 37"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m28 4-7 40"/>
              </svg>
            </div>
            <!-- 代码块 -->
            <div class="bytemd-toolbar-icon" title="代码块" @click="insertCodeBlock">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4c-2 0-5 1-5 5v9c0 3-5 5-5 5s5 2 5 5v11c0 4 3 5 5 5M32 4c2 0 5 1 5 5v9c0 3 5 5 5 5s-5 2-5 5v11c0 4-3 5-5 5"/>
              </svg>
            </div>
            <!-- 无序列表 -->
            <div class="bytemd-toolbar-icon" title="无序列表" @click="insertUnorderedList">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M9 42a4 4 0 1 0 0-8 4 4 0 0 0 0 8ZM9 14a4 4 0 1 0 0-8 4 4 0 0 0 0 8ZM9 28a4 4 0 1 0 0-8 4 4 0 0 0 0 8Z"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 24h22M21 38h22M21 10h22"/>
              </svg>
            </div>
            <!-- 有序列表 -->
            <div class="bytemd-toolbar-icon" title="有序列表" @click="insertOrderedList">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 4v9M12 13H6M12 27H6M6 20s3-3 5 0-5 7-5 7M6 34.5s2-3 5-1 0 4.5 0 4.5 3 2.5 0 4.5-5-1-5-1M11 38H9M9 4 6 6M21 24h22M21 38h22M21 10h22"/>
              </svg>
            </div>
            <!-- 删除线 -->
            <div class="bytemd-toolbar-icon" title="删除线" @click="insertStrikethrough">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 24h38M24 24c16 6 10 20 0 20s-12-8-12-8M36 12s-3-8-12-8-12.564 7.6-8.39 14"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36s4 8 12 8 12.564-7.6 8.39-14"/>
              </svg>
            </div>
            <!-- 表格 -->
            <div class="bytemd-toolbar-icon" title="表格" @click="insertTable">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-width="4" d="M39.3 6H8.7A2.7 2.7 0 0 0 6 8.7v30.6A2.7 2.7 0 0 0 8.7 42h30.6a2.7 2.7 0 0 0 2.7-2.7V8.7A2.7 2.7 0 0 0 39.3 6Z"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M18 6v36M30 6v36M6 18h36M6 30h36"/>
              </svg>
            </div>
            <!-- 任务列表 -->
            <div class="bytemd-toolbar-icon" title="任务列表" @click="insertTaskList">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-path="url(#a)">
                  <path d="M42 20v19a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h21"/>
                  <path d="m16 20 10 8L41 7"/>
                </g>
                <defs>
                  <clipPath id="a">
                    <path fill="currentColor" d="M0 0h48v48H0z"/>
                  </clipPath>
                </defs>
              </svg>
            </div>
            <!-- Mermaid图表 -->
            <div class="bytemd-toolbar-item relative">
              <div class="bytemd-toolbar-icon" title="Mermaid图表" @mouseenter="showMermaidDropdown = true">
                <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                  <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M17 6h14v9H17zM6 33h14v9H6zM28 33h14v9H28z"/>
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 16v8M13 33v-9h22v9"/>
                </svg>
              </div>
              <!-- Mermaid图表类型下拉菜单 -->
              <div v-if="showMermaidDropdown" class="mermaid-dropdown" 
                   @mouseenter="showMermaidDropdown = true"
                   @mouseleave="showMermaidDropdown = false" 
                   @mouseover.stop>
                <div class="mermaid-option" @click.stop="insertMermaidChart('flowchart')">流程图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('sequence')">时序图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('class')">类图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('state')">状态图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('er')">关系图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('journey')">旅程图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('gantt')">甘特图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('pie')">饼状图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('mindmap')">思维导图</div>
                <div class="mermaid-option" @click.stop="insertMermaidChart('timeline')">时间轴</div>
              </div>
            </div>
          </div>
          
          <div class="bytemd-toolbar-right">
            <!-- 目录 -->
            <div class="bytemd-toolbar-icon" :class="{ 'bytemd-toolbar-icon-active': showToc }" title="目录" @click="toggleToc">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26 24H14M34 15H14M32 33H14"/>
              </svg>
            </div>
            <!-- 帮助 -->
            <div class="bytemd-toolbar-icon" :class="{ 'bytemd-toolbar-icon-active': showHelp }" title="帮助" @click="toggleHelp">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 28.625v-4a6 6 0 1 0-6-6"/>
                <path fill="currentColor" fill-rule="evenodd" d="M24 37.625a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/>
              </svg>
            </div>
            <!-- 仅编辑区 -->
            <div class="bytemd-toolbar-icon" :class="{ 'bytemd-toolbar-icon-active': editorMode === 'edit' }" title="仅编辑区" @click="switchMode('edit')">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <rect width="28" height="36" x="6" y="6" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 6v36"/>
              </svg>
            </div>
            <!-- 仅预览区 -->
            <div class="bytemd-toolbar-icon" :class="{ 'bytemd-toolbar-icon-active': editorMode === 'preview' }" title="仅预览区" @click="switchMode('preview')">
              <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                <rect width="28" height="36" x="14" y="6" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 6v36"/>
              </svg>
            </div>
          </div>
        </div>

        <!-- 编辑器主体 -->
        <div class="bytemd-body">
          <!-- 编辑模式 -->
          <div 
            v-if="editorMode === 'edit'" 
            class="bytemd-editor"
            :style="{ width: showToc || showHelp ? 'calc(100% - 280px)' : '100%' }"
            @dragover.prevent="handleDragOver"
            @dragenter.prevent="handleDragEnter"
            @dragleave.prevent="handleDragLeave"
            @drop.prevent="handleImageDrop"
            :class="{ 'drag-over': isDragOver }"
          >
            <div class="editor-content-wrapper">
              <textarea 
                v-model="content" 
                class="markdown-textarea"
                placeholder="开始写作..."
                @input="handleContentChange"
                @dragover.prevent
                @dragenter.prevent
                @dragleave.prevent
                @drop.prevent="handleImageDrop"
              ></textarea>
            </div>
            <div v-if="isDragOver" class="drag-overlay" @click.self="closeDragOverlay">
              <div class="drag-overlay-content">
                <button class="drag-overlay-close" @click.stop="closeDragOverlay" aria-label="关闭">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </button>
                <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" class="upload-icon">
                  <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M17 8l-5-5-5 5M12 3v12"/>
                </svg>
                <p>拖拽图片到这里上传</p>
                <button class="select-image-btn" @click="openFileSelector">点击选择图片</button>
                <p class="drag-overlay-hint">支持 JPG、PNG、GIF 等格式，最大 10MB</p>
              </div>
            </div>
          </div>
          
          <!-- 预览模式 -->
          <div v-else-if="editorMode === 'preview'" class="bytemd-preview" :style="{ width: showToc || showHelp ? 'calc(100% - 280px)' : '100%' }">
            <div class="preview-content-wrapper">
              <div class="markdown-body" v-html="renderedContent"></div>
            </div>
          </div>
          
          <!-- 分屏模式（默认） -->
          <div v-else class="bytemd-split">
            <div 
              class="bytemd-editor" 
              :style="{ width: showToc || showHelp ? 'calc(50% - 140px)' : '50%' }"
              @dragover.prevent="handleDragOver"
              @dragenter.prevent="handleDragEnter"
              @dragleave.prevent="handleDragLeave"
              @drop.prevent="handleImageDrop"
              :class="{ 'drag-over': isDragOver }"
            >
              <div class="editor-content-wrapper">
                <textarea 
                  v-model="content" 
                  class="markdown-textarea"
                  placeholder="开始写作..."
                  @input="handleContentChange"
                  @dragover.prevent
                  @dragenter.prevent
                  @dragleave.prevent
                  @drop.prevent="handleImageDrop"
                ></textarea>
              </div>
              <div v-if="isDragOver" class="drag-overlay" @click.self="closeDragOverlay">
                <div class="drag-overlay-content">
                  <button class="drag-overlay-close" @click.stop="closeDragOverlay" aria-label="关闭">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <line x1="18" y1="6" x2="6" y2="18"></line>
                      <line x1="6" y1="6" x2="18" y2="18"></line>
                    </svg>
                  </button>
                  <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" class="upload-icon">
                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M17 8l-5-5-5 5M12 3v12"/>
                  </svg>
                  <p>拖拽图片到这里上传</p>
                  <button class="select-image-btn" @click="openFileSelector">点击选择图片</button>
                  <p class="drag-overlay-hint">支持 JPG、PNG、GIF 等格式，最大 10MB</p>
                </div>
              </div>
            </div>
            <div class="bytemd-preview" :style="{ width: showToc || showHelp ? 'calc(50% - 140px)' : '50%' }">
              <div class="preview-content-wrapper">
                <div class="markdown-body" v-html="renderedContent"></div>
              </div>
            </div>
          </div>
          
          <!-- 侧边栏 -->
          <div v-if="showToc || showHelp" class="bytemd-sidebar">
            
            <!-- 帮助面板 -->
            <div v-if="showHelp" class="bytemd-help">
              <div>
                <h2>帮助</h2>
              </div>
              
              <h2>Markdown 语法</h2>
              <ul>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 8v32M25 8v32M6 24h19M34.226 24 39 19.017V40"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">一级标题</div>
                  <div class="bytemd-help-content"><code># 标题</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 8v32M15 8v32M6 24h9M34.226 24 39 19.017V40"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">二级标题</div>
                  <div class="bytemd-help-content"><code>## 标题</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24c5.506 0 9.969-4.477 9.969-10S29.506 4 24 4H11v20h13ZM28.031 44C33.537 44 38 39.523 38 34s-4.463-10-9.969-10H11v20h17.031Z" clip-rule="evenodd"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">粗体</div>
                  <div class="bytemd-help-content"><code>**粗体文本**</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 6h16M12 42h16M29 5.952 19 42"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">斜体</div>
                  <div class="bytemd-help-content"><code>*斜体文本*</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path fill="currentColor" fill-rule="evenodd" d="M18.853 9.116C11.323 13.952 7.14 19.58 6.303 26.003 5 36 13.94 40.893 18.47 36.497 23 32.1 20.285 26.52 17.005 24.994c-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.016-7.69 9.116-10.322a.749.749 0 0 0 .114-1.02L20.285 9.3c-.44-.572-.862-.55-1.432-.185ZM38.679 9.116c-7.53 4.836-11.714 10.465-12.55 16.887-1.303 9.997 7.637 14.89 12.167 10.494 4.53-4.397 1.815-9.977-1.466-11.503-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.017-7.69 9.117-10.322a.749.749 0 0 0 .113-1.02L40.11 9.3c-.44-.572-.862-.55-1.431-.185Z" clip-rule="evenodd"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">引用</div>
                  <div class="bytemd-help-content"><code>&gt; 引用文本</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26.24 16.373-9.14-9.14c-2.661-2.661-7.035-2.603-9.768.131-2.734 2.734-2.793 7.107-.131 9.768l7.935 7.936M32.903 23.003l7.935 7.935c2.661 2.662 2.603 7.035-.13 9.769-2.735 2.734-7.108 2.792-9.77.13l-9.14-9.14"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26.11 26.142c2.733-2.734 2.792-7.108.13-9.769M21.799 21.798c-2.734 2.734-2.792 7.108-.131 9.769"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">链接</div>
                  <div class="bytemd-help-content"><code>[链接文本](URL)</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 10a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V10Z" clip-rule="evenodd"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14.5 18a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>
                      <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m15 24 5 4 6-7 17 13v4a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-4l10-10Z"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">图片</div>
                  <div class="bytemd-help-content"><code>![图片描述](图片URL)</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 13 4 25.432 16 37M32 13l12 12.432L32 37"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m28 4-7 40"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">行内代码</div>
                  <div class="bytemd-help-content"><code>行内代码`</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4c-2 0-5 1-5 5v9c0 3-5 5-5 5s5 2 5 5v11c0 4 3 5 5 5M32 4c2 0 5 1 5 5v9c0 3 5 5 5 5s-5 2-5 5v11c0 4-3 5-5 5"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">代码块</div>
                  <div class="bytemd-help-content"><code>```语言
代码块
```</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M9 42a4 4 0 1 0 0-8 4 4 0 0 0 0 8ZM9 14a4 4 0 1 0 0-8 4 4 0 0 0 0 8ZM9 28a4 4 0 1 0 0-8 4 4 0 0 0 0 8Z"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 24h22M21 38h22M21 10h22"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">无序列表</div>
                  <div class="bytemd-help-content"><code>- 列表项1
- 列表项2</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 4v9M12 13H6M12 27H6M6 20s3-3 5 0-5 7-5 7M6 34.5s2-3 5-1 0 4.5 0 4.5 3 2.5 0 4.5-5-1-5-1M11 38H9M9 4 6 6M21 24h22M21 38h22M21 10h22"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">有序列表</div>
                  <div class="bytemd-help-content"><code>1. 列表项1
2. 列表项2</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 24h38M24 24c16 6 10 20 0 20s-12-8-12-8M36 12s-3-8-12-8-12.564 7.6-8.39 14"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36s4 8 12 8 12.564-7.6 8.39-14"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">删除线</div>
                  <div class="bytemd-help-content"><code>~~删除文本~~</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-width="4" d="M39.3 6H8.7A2.7 2.7 0 0 0 6 8.7v30.6A2.7 2.7 0 0 0 8.7 42h30.6a2.7 2.7 0 0 0 2.7-2.7V8.7A2.7 2.7 0 0 0 39.3 6Z"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M18 6v36M30 6v36M6 18h36M6 30h36"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">表格</div>
                  <div class="bytemd-help-content"><code>| 表头1 | 表头2 |
| --- | --- |
| 内容1 | 内容2 |</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-path="url(#a)">
                        <path d="M42 20v19a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h21"/>
                        <path d="m16 20 10 8L41 7"/>
                      </g>
                      <defs>
                        <clipPath id="a">
                          <path fill="currentColor" d="M0 0h48v48H0z"/>
                        </clipPath>
                      </defs>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">任务列表</div>
                  <div class="bytemd-help-content"><code>- [x] 已完成任务
- [ ] 未完成任务</code></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M17 6h14v9H17zM6 33h14v9H6zM28 33h14v9H28z"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 16v8M13 33v-9h22v9"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">Mermaid图表</div>
                  <div class="bytemd-help-content"><code>```mermaid
flowchart TD
    A --> B
```</code></div>
                </li>
              </ul>
              
              <h2>快捷键</h2>
              <ul>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 24c5.506 0 9.969-4.477 9.969-10S29.506 4 24 4H11v20h13ZM28.031 44C33.537 44 38 39.523 38 34s-4.463-10-9.969-10H11v20h17.031Z" clip-rule="evenodd"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">粗体</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-B</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 6h16M12 42h16M29 5.952 19 42"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">斜体</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-I</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path fill="currentColor" fill-rule="evenodd" d="M18.853 9.116C11.323 13.952 7.14 19.58 6.303 26.003 5 36 13.94 40.893 18.47 36.497 23 32.1 20.285 26.52 17.005 24.994c-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.016-7.69 9.116-10.322a.749.749 0 0 0 .114-1.02L20.285 9.3c-.44-.572-.862-.55-1.432-.185ZM38.679 9.116c-7.53 4.836-11.714 10.465-12.55 16.887-1.303 9.997 7.637 14.89 12.167 10.494 4.53-4.397 1.815-9.977-1.466-11.503-3.28-1.525-5.286-.994-4.936-3.033.35-2.038 5.017-7.69 9.117-10.322a.749.749 0 0 0 .113-1.02L40.11 9.3c-.44-.572-.862-.55-1.431-.185Z" clip-rule="evenodd"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">引用</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-Q</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26.24 16.373-9.14-9.14c-2.661-2.661-7.035-2.603-9.768.131-2.734 2.734-2.793 7.107-.131 9.768l7.935 7.936M32.903 23.003l7.935 7.935c2.661 2.662 2.603 7.035-.13 9.769-2.735 2.734-7.108 2.792-9.77.13l-9.14-9.14"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26.11 26.142c2.733-2.734 2.792-7.108.13-9.769M21.799 21.798c-2.734 2.734-2.792 7.108-.131 9.769"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">链接</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-K</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4c-2 0-5 1-5 5v9c0 3-5 5-5 5s5 2 5 5v11c0 4 3 5 5 5M32 4c2 0 5 1 5 5v9c0 3 5 5 5 5s-5 2-5 5v11c0 4-3 5-5 5"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">代码块</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-Alt-C</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 13 4 25.432 16 37M32 13l12 12.432L32 37"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m28 4-7 40"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">行内代码</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-Shift-C</kbd></div>
                </li>
                <li>
                  <div class="bytemd-help-icon">
                    <svg width="1em" height="1em" fill="none" viewBox="0 0 48 48">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/>
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26 24H14M34 15H14M32 33H14"/>
                    </svg>
                  </div>
                  <div class="bytemd-help-title">显示/隐藏目录</div>
                  <div class="bytemd-help-content"><kbd>Ctrl-Alt-T</kbd></div>
                </li>
              </ul>
            </div>
            
            <!-- 目录 -->
            <div v-if="showToc" class="bytemd-toc">
              <h2>目录</h2>
              <ul>
                <li v-for="item in tocItems" :key="item.id" :class="`toc-level-${item.level}`">
                  <a @click="jumpToHeading(item.id)">{{ item.text }}</a>
                </li>
              </ul>
            </div>
          </div>
        </div>

      </div>
    </div>
    
    <!-- 状态栏 -->
    <div class="bytemd-status">
      <div class="bytemd-status-left">
        <span>字符数: <strong>{{ characterCount }}</strong></span>
        <span>行数: <strong>{{ lineCount }}</strong></span>
      </div>
           <div class="bytemd-status-right">
             <label>
               <input type="checkbox" v-model="syncScroll" :disabled="editorMode !== 'split'"> 同步滚动
             </label>
             <span @click="scrollToTop" style="cursor: pointer;">回到顶部</span>
           </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed, onMounted, onUnmounted, nextTick } from 'vue'
import blogApi from '../services/api.js'
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

// Props
const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'update:title', 'save'])

// 响应式数据
const title = ref(props.title || '')
const content = ref(props.modelValue || '')
const tocItems = ref([])
const editorMode = ref('split') // edit, preview, split
const showToc = ref(false)
const showHelp = ref(false)
const isDragOver = ref(false)
const showHeadingDropdown = ref(false)
const showMermaidDropdown = ref(false) // Mermaid图表下拉菜单显示控制
const syncScroll = ref(true) // 同步滚动功能

// 计算属性
const characterCount = computed(() => {
  return content.value.length
})

const lineCount = computed(() => {
  return content.value.split('\n').length
})

// 创建MarkdownIt实例并配置插件
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
  breaks: true, // 支持换行符转换为<br>
  highlight: function(code, lang) {
    if (lang && hljs.getLanguage(lang)) {
      try {
        return hljs.highlight(code, { language: lang }).value
      } catch (err) {
        // 高亮显示失败，返回原代码
        return code
      }
    }
    try {
      return hljs.highlightAuto(code).value
    } catch (err) {
      // 自动高亮失败，返回原代码
      return code
    }
  }
})

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

// 注意：已移除emoji插件，博客不再支持表情符号功能

// 注意：markdown-it默认已支持基本表格语法，无需额外插件
// 表格语法: | 表头1 | 表头2 |\n| --- | --- |\n| 单元格1 | 单元格2 |

// 为mermaid图表添加自定义规则处理
md.renderer.rules.fence = function(tokens, idx, options, env, self) {
  const token = tokens[idx];
  const info = token.info ? token.info.trim() : '';
  
  if (info === 'mermaid') {
    // 对于mermaid图表，直接创建带有mermaid类的div，让mermaid库自动识别
    const code = token.content.trim();
    const id = `mermaid-${Date.now()}-${Math.floor(Math.random() * 1000)}`;
    return `<div class="mermaid-chart-container" style="margin: 16px 0; padding: 16px; background-color: #f8f9fa; border-radius: 4px;">
      <div id="${id}" class="mermaid">${md.utils.escapeHtml(code)}</div>
    </div>`;
  }
  
  // 对于其他代码块，使用默认渲染器
  return self.renderToken(tokens, idx, options);
}

// 自定义图片渲染器处理图片路径
const defaultImageRenderer = md.renderer.rules.image_open || function(tokens, idx, options, env, self) {
  return self.renderToken(tokens, idx, options)
}

md.renderer.rules.image_open = function(tokens, idx, options, env, self) {
  const token = tokens[idx]
  let srcAttr = token.attrGet('src')
  
  if (srcAttr && !srcAttr.startsWith('http') && !srcAttr.startsWith('/')) {
    token.attrSet('src', `/api/uploads/${srcAttr}`)
  }
  
  return defaultImageRenderer(tokens, idx, options, env, self)
}

// 初始化mermaid
if (typeof window !== 'undefined') {
  // 确保只初始化一次
  if (!window.mermaidInitialized) {
    window.mermaidInitialized = true
    mermaid.initialize({
      startOnLoad: true, // 启用自动加载渲染
      securityLevel: 'loose', // 允许更灵活的图表内容
      maxTextSize: 20000 // 增加最大文本大小以支持复杂图表
    })
  }
}

const renderedContent = computed(() => {
  if (!content.value) return ''
  
  // 使用markdown-it渲染Markdown内容
  const html = md.render(content.value)
  
  // 渲染完成后确保mermaid图表被正确渲染
  nextTick(() => {
    if (typeof window !== 'undefined' && window.mermaidInitialized) {
      // 延迟执行，确保DOM已完全更新
      setTimeout(() => {
        try {
          // 查找所有带有mermaid类的元素并强制渲染
          const mermaidElements = document.querySelectorAll('.mermaid')
          if (mermaidElements.length > 0) {
            // 调用mermaid的run方法渲染所有图表
            mermaid.run().catch(() => {
              // 静默处理错误，避免控制台污染
            })
          }
        } catch (e) {
          // 静默处理错误，避免控制台污染
        }
      }, 300) // 稍长一点的延迟，确保DOM完全更新
    }
  })
  
  return html
})

// 方法
// 保存上一次的内容，用于检测图片删除
let previousContent = ''

// 从文本中提取所有图片URL的函数
const extractImageUrls = (text) => {
  const regex = /!\[.*?\]\((\/uploads\/images\/[^)]+)\)/g
  const urls = []
  let match
  while ((match = regex.exec(text)) !== null) {
    urls.push(match[1])
  }
  return urls
}

// 删除服务器上的图片文件
const deleteServerImage = async (imageUrl) => {
  try {
    // 从URL中提取文件名
    const parts = imageUrl.split('/')
    const filename = parts[parts.length - 1]
    
    // 检查是否有token
    const token = localStorage.getItem('token')
    if (!token) {
      return
    }
    
    // 使用API服务调用删除接口
    await blogApi.deleteImage(filename)
  } catch (error) {
    // 删除图片失败，静默处理
  }
}

const handleContentChange = () => {
  // 保存当前内容以便下次比较
  const currentContent = content.value
  
  // 检测并删除被移除的图片
  if (previousContent) {
    const previousImages = extractImageUrls(previousContent)
    const currentImages = extractImageUrls(currentContent)
    
    // 找出被删除的图片URL
    const deletedImages = previousImages.filter(url => !currentImages.includes(url))
    
    // 为每个被删除的图片调用删除API
    deletedImages.forEach(imageUrl => {
      deleteServerImage(imageUrl)
    })
  }
  
  // 更新上一次的内容
  previousContent = currentContent
  
  emit('update:modelValue', content.value)
  generateToc()
}

const handleTitleChange = () => {
  emit('update:title', title.value)
}

const generateToc = () => {
  const lines = content.value.split('\n')
  const toc = []
  
  lines.forEach((line, index) => {
    const match = line.match(/^(#{1,6})\s+(.+)/)
    if (match) {
      const level = match[1].length
      const text = match[2].trim()
      const id = `heading-${index}`
      
      toc.push({
        id,
        level,
        text,
        line: index
      })
    }
  })
  
  tocItems.value = toc
}

// 全局点击事件处理函数，用于关闭下拉菜单
const handleGlobalClick = (event) => {
  // 简单直接的逻辑：如果点击的不是下拉菜单本身，就关闭所有下拉菜单
  // 这样可以确保点击任何地方（除了下拉菜单内部）都会关闭菜单
  if (!event.target.closest('.heading-dropdown')) {
    showHeadingDropdown.value = false
  }
  
  if (!event.target.closest('.mermaid-dropdown')) {
    showMermaidDropdown.value = false
  }
}

// 全局鼠标移动事件处理函数，用于关闭下拉菜单
const handleGlobalMouseMove = (event) => {
  // 获取当前鼠标位置的元素
  const targetElement = event.target;
  
  // 检查是否在工具栏图标或下拉菜单内部
  const isInHeadingRelated = targetElement.closest('.bytemd-toolbar-item.relative') && 
                            (targetElement.closest('.heading-dropdown') || 
                             targetElement.closest('.bytemd-toolbar-icon') && 
                             targetElement.closest('.bytemd-toolbar-item').querySelector('.heading-dropdown'));
  
  const isInMermaidRelated = targetElement.closest('.bytemd-toolbar-item.relative') && 
                            (targetElement.closest('.mermaid-dropdown') || 
                             targetElement.closest('.bytemd-toolbar-icon') && 
                             targetElement.closest('.bytemd-toolbar-item').querySelector('.mermaid-dropdown'));
  
  // 如果鼠标移出了标题相关区域，关闭标题菜单
  if (!isInHeadingRelated && showHeadingDropdown.value) {
    showHeadingDropdown.value = false;
  }
  
  // 如果鼠标移出了mermaid相关区域，关闭mermaid菜单
  if (!isInMermaidRelated && showMermaidDropdown.value) {
    showMermaidDropdown.value = false;
  }
}

// 生命周期钩子
onMounted(() => {
  // 添加全局点击事件监听器
  document.addEventListener('click', handleGlobalClick);
  // 添加全局鼠标移动事件监听器
  document.addEventListener('mousemove', handleGlobalMouseMove);
})

onUnmounted(() => {
  // 清理全局点击事件监听器
  document.removeEventListener('click', handleGlobalClick);
  // 清理全局鼠标移动事件监听器
  document.removeEventListener('mousemove', handleGlobalMouseMove);
})

const jumpToHeading = (id) => {
  // 实现跳转到标题的功能
  // 预留实现空间
}


// 回到顶部
const scrollToTop = () => {
  try {
    // 滚动页面到顶部
    window.scrollTo({ top: 0, behavior: 'smooth' })
    
    // 查找并滚动所有可能的滚动容器
    const scrollableSelectors = [
      '.bytemd-editor',
      '.bytemd-preview', 
      '.markdown-body',
      '.bytemd-body',
      '.juejin-editor',
      '.bytemd',
      '.juejin-main',
      '.editor-panel',
      '.preview-panel',
      '.CodeMirror',
      '.CodeMirror-scroll',
      '.CodeMirror-sizer',
      '.CodeMirror-lines',
      '.CodeMirror-code',
      '.CodeMirror-wrap',
      '.CodeMirror-scrollbar',
      '.CodeMirror-vscrollbar',
      '.CodeMirror-hscrollbar'
    ]
    
    scrollableSelectors.forEach(selector => {
      const elements = document.querySelectorAll(selector)
      elements.forEach(element => {
        if (element) {
          element.scrollTop = 0
          if (element.scrollTo) {
            element.scrollTo({ top: 0, behavior: 'smooth' })
          }
        }
      })
    })
    
    // 多次尝试滚动确保成功
    for (let i = 0; i < 3; i++) {
      setTimeout(() => {
        const editorElement = document.querySelector('.bytemd-editor')
        if (editorElement) {
          editorElement.scrollTop = 0
        }
        
        const codeMirrorElement = document.querySelector('.CodeMirror')
        if (codeMirrorElement) {
          codeMirrorElement.scrollTop = 0
        }
        
        const codeMirrorScroll = document.querySelector('.CodeMirror-scroll')
        if (codeMirrorScroll) {
          codeMirrorScroll.scrollTop = 0
        }
      }, i * 50)
    }
    
    // 最终强制滚动所有可滚动元素
    setTimeout(() => {
      const allElements = document.querySelectorAll('*')
      allElements.forEach(element => {
        if (element.scrollHeight > element.clientHeight) {
          element.scrollTop = 0
        }
      })
    }, 200)
  } catch (error) {
    // 静默失败，避免控制台错误
  }
}

// 同步滚动功能
let syncScrollCleanup = null
let isScrolling = false
let syncScrollSetupAttempts = 0

const setupSyncScroll = () => {
  if (!syncScroll.value || editorMode.value !== 'split') {
    return
  }
  
  // 等待编辑器完全加载
  setTimeout(() => {
    // 再次检查syncScroll状态，确保在延迟期间没有被禁用
    if (!syncScroll.value) {
      return
    }
    
    // 查找编辑区滚动容器 - 优先使用外层容器
    let editorScrollElement = document.querySelector('.bytemd-editor')
    
    // 如果外层容器没有滚动，查找内层滚动元素
    if (!editorScrollElement || editorScrollElement.scrollHeight <= editorScrollElement.clientHeight) {
      // 尝试查找CodeMirror滚动容器
      editorScrollElement = document.querySelector('.bytemd-editor .CodeMirror-scroll')
      if (!editorScrollElement) {
        editorScrollElement = document.querySelector('.bytemd-editor .CodeMirror')
      }
      if (!editorScrollElement) {
        // 查找其他可能的滚动元素
        const bytemdEditor = document.querySelector('.bytemd-editor')
        if (bytemdEditor) {
          const scrollableElements = bytemdEditor.querySelectorAll('*')
          for (const el of scrollableElements) {
            if (el.scrollHeight > el.clientHeight) {
              editorScrollElement = el
              break
            }
          }
        }
      }
    }
    
    // 查找预览区滚动容器 - 优先使用外层容器
    let previewScrollElement = document.querySelector('.bytemd-preview');
    
    // 如果外层容器没有滚动，查找内层滚动元素
    if (!previewScrollElement || previewScrollElement.scrollHeight <= previewScrollElement.clientHeight) {
      previewScrollElement = document.querySelector('.bytemd-preview .markdown-body') ||
                             document.querySelector('.bytemd-preview');
    }
    
    if (!editorScrollElement || !previewScrollElement) {
      // 只有在syncScroll仍为true时才重试
      if (syncScroll.value && syncScrollSetupAttempts < 5) {
        syncScrollSetupAttempts++;
        setTimeout(setupSyncScroll, 1000);
      }
      return;
    }
      
    // 清理之前的监听器
    if (syncScrollCleanup) {
      syncScrollCleanup();
    }
    
    // 滚动事件处理
    function handleEditorScroll() {
      if (isScrolling || !syncScroll.value) return;
      isScrolling = true;
      
      try {
        if (!editorScrollElement || !previewScrollElement) return;
        
        const editorScrollHeight = editorScrollElement.scrollHeight - editorScrollElement.clientHeight;
        const previewScrollHeight = previewScrollElement.scrollHeight - previewScrollElement.clientHeight;
        
        if (editorScrollHeight > 0 && previewScrollHeight > 0) {
          const ratio = Math.min(1, Math.max(0, editorScrollElement.scrollTop / editorScrollHeight));
          const targetScroll = ratio * previewScrollHeight;
          
          window.requestAnimationFrame(function() {
            if (syncScroll.value && previewScrollElement) {
              previewScrollElement.scrollTop = targetScroll;
            }
          });
        }
      } catch (err) {
        // 静默失败
      }
      
      setTimeout(function() {
        isScrolling = false;
      }, 100);
    }
    
    function handlePreviewScroll() {
      if (isScrolling || !syncScroll.value) return;
      isScrolling = true;
      
      try {
        if (!editorScrollElement || !previewScrollElement) return;
        
        const editorScrollHeight = editorScrollElement.scrollHeight - editorScrollElement.clientHeight;
        const previewScrollHeight = previewScrollElement.scrollHeight - previewScrollElement.clientHeight;
        
        if (previewScrollHeight > 0 && editorScrollHeight > 0) {
          const ratio = Math.min(1, Math.max(0, previewScrollElement.scrollTop / previewScrollHeight));
          const targetScroll = ratio * editorScrollHeight;
          
          window.requestAnimationFrame(function() {
            if (syncScroll.value && editorScrollElement) {
              editorScrollElement.scrollTop = targetScroll;
            }
          });
        }
      } catch (err) {
        // 静默失败
      }
      
      setTimeout(function() {
        isScrolling = false;
      }, 100);
    }
    
    // 添加事件监听器
    editorScrollElement.addEventListener('scroll', handleEditorScroll, { passive: true });
    previewScrollElement.addEventListener('scroll', handlePreviewScroll, { passive: true });
    
    // 设置清理函数
    syncScrollCleanup = function() {
      editorScrollElement.removeEventListener('scroll', handleEditorScroll);
      previewScrollElement.removeEventListener('scroll', handlePreviewScroll);
      isScrolling = false;
    };
  }, 1000);
}

// 监听同步滚动状态变化
watch(syncScroll, (newValue) => {
  // 清理之前的监听器
  if (syncScrollCleanup) {
    syncScrollCleanup()
    syncScrollCleanup = null
  }
  
  // 重置尝试次数和状态
  syncScrollSetupAttempts = 0
  isScrolling = false
  
  // 当禁用同步滚动时，额外执行一次全局清理，确保没有残留的滚动行为
  if (!newValue) {
    // 手动移除可能存在的滚动监听器
    const editorElements = [
      document.querySelector('.bytemd-editor'),
      document.querySelector('.bytemd-editor .CodeMirror-scroll'),
      document.querySelector('.bytemd-editor .CodeMirror')
    ].filter(Boolean)
    
    const previewElements = [
      document.querySelector('.bytemd-preview'),
      document.querySelector('.bytemd-preview .markdown-body')
    ].filter(Boolean)
    
    // 为每个元素移除所有scroll事件监听器（通过克隆和替换元素的方式）
    editorElements.concat(previewElements).forEach(function(el) {
      const clone = el.cloneNode(true);
      if (el.parentNode) {
        el.parentNode.replaceChild(clone, el);
      }
    });
  } else if (editorMode.value === 'split') { // 只有在双屏模式下才设置同步滚动
    setupSyncScroll()
  }
})

// 模式切换
const switchMode = (mode) => {
  if (mode === 'edit') {
    // 如果当前是编辑模式，点击后取消选中，回到分屏模式
    if (editorMode.value === 'edit') {
      editorMode.value = 'split'
    } else {
      editorMode.value = 'edit'
    }
  } else if (mode === 'preview') {
    // 如果当前是预览模式，点击后取消选中，回到分屏模式
    if (editorMode.value === 'preview') {
      editorMode.value = 'split'
    } else {
      editorMode.value = 'preview'
    }
  } else {
    editorMode.value = mode
  }
  
  // 模式切换后，如果离开分屏模式，清理同步滚动监听器
  if (editorMode.value !== 'split' && syncScrollCleanup) {
    syncScrollCleanup()
    syncScrollCleanup = null
    isScrolling = false
  }
  
  // 模式切换后，如果进入分屏模式且同步滚动开启，重新设置同步滚动
  nextTick(() => {
    if (syncScroll.value && editorMode.value === 'split') {
      setTimeout(() => {
        setupSyncScroll()
      }, 200)
    }
  })
}

// 切换目录显示
const toggleToc = () => {
  if (showToc.value) {
    showToc.value = false
  } else {
    showToc.value = true
    showHelp.value = false // 关闭帮助面板
    generateToc()
  }
}

// 切换帮助显示
const toggleHelp = () => {
  if (showHelp.value) {
    showHelp.value = false
  } else {
    showHelp.value = true
    showToc.value = false // 关闭目录面板
  }
}

// 处理发布
const handlePublish = () => {
  if (!title.value.trim()) {
    alert('请输入文章标题')
    return
  }
  if (!content.value.trim()) {
    alert('请输入文章内容')
    return
  }
  
  // 触发保存事件，由父组件处理发布逻辑
  emit('save')
}

// 处理草稿箱
const handleDraft = () => {
  if (!title.value.trim()) {
    alert('请输入文章标题')
    return
  }
  if (!content.value.trim()) {
    alert('请输入文章内容')
    return
  }
  
  // 触发保存事件，由父组件处理草稿逻辑
  emit('save')
}


// 关闭侧边栏
const closeSidebar = () => {
  showToc.value = false
  showHelp.value = false
}

// 工具栏图标点击事件处理函数
const insertHeading = () => {
  // 默认使用一级标题
  insertHeadingLevel(1)
}

// 插入指定级别的标题
const insertHeadingLevel = (level) => {
  // 确保级别在有效范围内
  level = Math.max(1, Math.min(6, level))
  
  const selectedText = getSelectedText()
  const hashSigns = '#'.repeat(level)
  
  if (selectedText) {
    // 如果有选中文本，直接包装为标题
    const heading = `${hashSigns} ${selectedText}\n\n`
    replaceSelectedText(heading)
  } else {
    // 如果没有选中文本，插入模板
    const heading = `${hashSigns} 标题${level}\n\n`
    insertTextAtCursor(heading)
    selectTextInEditor(`标题${level}`)
  }
  
  // 关闭下拉菜单
  showHeadingDropdown.value = false
}

const insertBold = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装
    const bold = `**${selectedText}**`
    replaceSelectedText(bold)
  } else {
    // 如果没有选中文本，插入模板
    const bold = `**粗体文本**`
    insertTextAtCursor(bold)
    selectTextInEditor(bold)
  }
}

const insertItalic = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装
    const italic = `*${selectedText}*`
    replaceSelectedText(italic)
  } else {
    // 如果没有选中文本，插入模板
    const italic = `*斜体文本*`
    insertTextAtCursor(italic)
    selectTextInEditor(italic)
  }
}

const insertQuote = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装
    const quote = `> ${selectedText}\n\n`
    replaceSelectedText(quote)
  } else {
    // 如果没有选中文本，插入模板
    const quote = `> 引用内容\n\n`
    insertTextAtCursor(quote)
    selectTextInEditor('引用内容')
  }
}

const insertLink = () => {
  const selectedText = getSelectedText()
  
  // 尝试使用CodeMirror方式
  const editor = document.querySelector('.bytemd-editor .CodeMirror')
  if (editor && editor.CodeMirror) {
    const cm = editor.CodeMirror
    const doc = cm.getDoc()
    
    if (selectedText) {
      // 有选中文本的情况
      const link = `[${selectedText}](链接地址)`
      doc.replaceSelection(link)
      
      // 获取当前光标位置
      const cursor = doc.getCursor()
      // 从后向前搜索"链接地址"
      const lineContent = doc.getLine(cursor.line)
      const linkAddrPos = lineContent.lastIndexOf("链接地址", cursor.ch)
      
      if (linkAddrPos !== -1) {
        const start = { line: cursor.line, ch: linkAddrPos }
        const end = { line: cursor.line, ch: linkAddrPos + 4 }
        doc.setSelection(start, end)
      }
    } else {
      // 没有选中文本的情况
      const link = `[链接文本](链接地址)`
      const cursor = doc.getCursor()
      doc.replaceRange(link, cursor)
      
      // 计算新光标位置（插入文本后）
      const newCursor = { line: cursor.line, ch: cursor.ch + link.length }
      // 选中"链接文本"
      const start = { line: cursor.line, ch: cursor.ch + 1 }
      const end = { line: cursor.line, ch: cursor.ch + 5 }
      doc.setSelection(start, end)
    }
    
    cm.focus()
  } else {
    // 尝试使用原生textarea方式
    const textarea = document.querySelector('.bytemd-editor .markdown-textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      const currentValue = textarea.value
      
      if (selectedText) {
        // 有选中文本的情况
        const link = `[${selectedText}](链接地址)`
        const newValue = currentValue.substring(0, start) + link + currentValue.substring(end)
        textarea.value = newValue
        content.value = newValue
        emit('update:modelValue', newValue)
        
        // 计算并选中链接地址
        const newCursorPos = start + link.length
        const linkAddrPos = newValue.lastIndexOf("链接地址", newCursorPos)
        if (linkAddrPos !== -1) {
          textarea.setSelectionRange(linkAddrPos, linkAddrPos + 4)
        }
      } else {
        // 没有选中文本的情况
        const link = `[链接文本](链接地址)`
        const newValue = currentValue.substring(0, start) + link + currentValue.substring(end)
        textarea.value = newValue
        content.value = newValue
        emit('update:modelValue', newValue)
        
        // 选中链接文本
        textarea.setSelectionRange(start + 1, start + 5)
      }
      
      textarea.focus()
    }
  }
}

// 创建Toast提示的公共函数
const createToast = (message, type = 'info') => {
  const toast = document.createElement('div')
  const colors = {
    info: '#17a2b8',
    success: '#28a745', 
    warning: '#ffc107',
    error: '#dc3545'
  }
  
  toast.style.cssText = `
    position: fixed;
    top: 20px;
    right: 20px;
    background: ${colors[type] || colors.info};
    color: white;
    padding: 12px 20px;
    border-radius: 6px;
    z-index: 10000;
    font-size: 14px;
  `
  toast.textContent = message
  document.body.appendChild(toast)
  return toast
}

// 移除Toast的公共函数
const removeToast = (toast) => {
  if (document.body.contains(toast)) {
    document.body.removeChild(toast)
  }
}

// 调整图片尺寸的函数
const resizeImage = (file, maxWidth = 800) => {
  return new Promise((resolve) => {
    const img = new Image()
    const canvas = document.createElement('canvas')
    const ctx = canvas.getContext('2d')
    
    img.onload = () => {
      // 如果图片宽度小于等于最大宽度，直接返回原文件
      if (img.width <= maxWidth) {
        resolve(file)
        return
      }
      
      // 计算新的尺寸
      const ratio = maxWidth / img.width
      const newWidth = maxWidth
      const newHeight = img.height * ratio
      
      // 设置canvas尺寸
      canvas.width = newWidth
      canvas.height = newHeight
      
      // 绘制调整后的图片
      ctx.drawImage(img, 0, 0, newWidth, newHeight)
      
      // 转换为Blob
      canvas.toBlob((blob) => {
        // 创建新的File对象，保持原文件名
        const resizedFile = new File([blob], file.name, {
          type: file.type,
          lastModified: Date.now()
        })
        resolve(resizedFile)
      }, file.type, 0.9) // 使用90%质量
    }
    
    img.onerror = () => {
      // 如果图片加载失败，返回原文件
      resolve(file)
    }
    
    img.src = URL.createObjectURL(file)
  })
}

// 上传图片的通用函数
const uploadImage = async (file) => {
  try {
    // 自动使用文件名作为图片描述
    const alt = file.name.replace(/\.[^/.]+$/, '')
    
    // 显示图片处理中提示
    const processingToast = createToast('正在调整图片尺寸...', 'info')
    
    // 调整图片尺寸以适应预览区宽度（800px）
    let resizedFile
    try {
      resizedFile = await resizeImage(file, 800)
    } catch (error) {
      resizedFile = file
    }
    
    // 移除处理中提示
    removeToast(processingToast)
    
    // 创建FormData上传文件
    const formData = new FormData()
    formData.append('image', resizedFile)
    
    // 显示上传中提示
    const uploadToast = createToast('正在上传图片...', 'info')
    
    // 获取认证token
    const token = localStorage.getItem('token')
    if (!token) {
      alert('请先登录后再上传图片')
      removeToast(uploadToast)
      return
    }
    
    const headers = {
      'Authorization': `Bearer ${token}`
    }
    
    // 上传文件
        const response = await fetch('http://localhost:8080/api/upload/image', {
          method: 'POST',
          body: formData,
          credentials: 'include',
          headers: headers
        })
        
        // 移除上传提示
        removeToast(uploadToast)
        
        if (response.ok) {
          const result = await response.json()
          // 处理后端返回的格式：{"data": {"url": "...", ...}}
          if (result.data && result.data.url) {
            // 插入图片Markdown
            const image = `![${alt}](${result.data.url})\n\n`
            insertTextAtCursor(image)
            
            // 显示成功提示
            const successToast = createToast('图片上传成功！', 'success')
            
            // 3秒后移除成功提示
            setTimeout(() => {
              removeToast(successToast)
            }, 3000)
          } else if (result.error) {
            alert('图片上传失败: ' + result.error)
          } else {
            alert('图片上传失败: 服务器返回格式不正确')
          }
        } else {
          alert('图片上传失败，请重试')
        }
  } catch (error) {
    // 发生错误时只显示用户友好提示
    alert('图片上传失败，请重试')
  }
}

const insertImage = () => {
  // 显示拖拽覆盖层
  isDragOver.value = true
  
  // 添加全局点击事件监听器来关闭拖拽覆盖层
  const handleGlobalClick = (event) => {
    // 检查点击是否在编辑器区域外
    const editorElement = document.querySelector('.bytemd-editor')
    if (editorElement && !editorElement.contains(event.target) && !event.target.closest('.bytemd-toolbar-icon[title="图片"]')) {
      isDragOver.value = false
      document.removeEventListener('click', handleGlobalClick)
    }
  }
  
  setTimeout(() => {
    document.addEventListener('click', handleGlobalClick)
  }, 0)
  
  // 拖拽覆盖层已显示
}

// 关闭拖拽覆盖层
const closeDragOverlay = () => {
  isDragOver.value = false
}

const openFileSelector = async () => {
  // 创建文件输入元素
  const fileInput = document.createElement('input')
  fileInput.type = 'file'
  fileInput.accept = 'image/*'
  fileInput.multiple = false
  fileInput.style.display = 'none'
  
  // 添加到页面并触发点击
  document.body.appendChild(fileInput)
  fileInput.click()
  
  // 监听文件选择
  fileInput.addEventListener('change', async (event) => {
    const file = event.target.files[0]
    if (file) {
      // 验证文件类型
      if (!file.type.startsWith('image/')) {
        alert('请选择图片文件')
        document.body.removeChild(fileInput)
        return
      }
      
      // 验证文件大小 (限制为10MB)
      if (file.size > 10 * 1024 * 1024) {
        alert('图片文件过大，请选择小于10MB的图片')
        document.body.removeChild(fileInput)
        return
      }
      
      // 关闭拖拽覆盖层
      isDragOver.value = false
      
      await uploadImage(file)
    }
    
    // 清理文件输入元素
    if (document.body.contains(fileInput)) {
      document.body.removeChild(fileInput)
    }
  })
  
  // 监听取消事件
  fileInput.addEventListener('cancel', () => {
    if (document.body.contains(fileInput)) {
      document.body.removeChild(fileInput)
    }
  })
  
  // 监听焦点丢失事件
  fileInput.addEventListener('blur', () => {
    setTimeout(() => {
      if (document.body.contains(fileInput)) {
        document.body.removeChild(fileInput)
      }
    }, 100)
  })
}

const insertInlineCode = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装
    const code = `\`${selectedText}\``
    replaceSelectedText(code)
  } else {
    // 如果没有选中文本，插入模板
    const code = `\`代码\``
    insertTextAtCursor(code)
    selectTextInEditor('代码')
  }
}

const insertCodeBlock = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装为代码块
    const codeBlock = `\`\`\`\n${selectedText}\n\`\`\`\n\n`
    replaceSelectedText(codeBlock)
  } else {
    // 如果没有选中文本，插入模板
    const codeBlock = `\`\`\`\n代码内容\n\`\`\`\n\n`
    insertTextAtCursor(codeBlock)
    selectTextInEditor('代码内容')
  }
}

const insertUnorderedList = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，转换为列表项
    const list = `- ${selectedText}\n`
    replaceSelectedText(list)
  } else {
    // 如果没有选中文本，插入模板
    const list = `- 列表项\n`
    insertTextAtCursor(list)
    selectTextInEditor('列表项')
  }
}

const insertOrderedList = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，转换为有序列表项
    const list = `1. ${selectedText}\n`
    replaceSelectedText(list)
  } else {
    // 如果没有选中文本，插入有序列表模板
    const list = `1. 列表项\n`
    insertTextAtCursor(list)
    selectTextInEditor('列表项')
  }
}

const insertStrikethrough = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，直接包装
    const strikethrough = `~~${selectedText}~~`
    replaceSelectedText(strikethrough)
  } else {
    // 如果没有选中文本，插入模板
    const strikethrough = `~~删除线文本~~`
    insertTextAtCursor(strikethrough)
    selectTextInEditor('删除线文本')
  }
}

const insertTable = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，尝试将其转换为表格格式
    // 这里简单地将每行文本作为表格的一行
    const lines = selectedText.split('\n').map(line => line.trim()).filter(line => line)
    if (lines.length > 0) {
      // 创建表头（使用第一行作为表头）
      let table = `| ${lines[0]} |\n`
      table += `| ${'-'.repeat(lines[0].length)} |\n`
      
      // 添加剩余行作为表格内容
      for (let i = 1; i < Math.min(lines.length, 5); i++) { // 最多显示前5行
        table += `| ${lines[i]} |\n`
      }
      
      table += '\n'
      replaceSelectedText(table)
      return
    }
  }
  
  // 如果没有选中文本或选中文本无法转换为表格，插入默认表格模板
  const table = `| 列1 | 列2 | 列3 |\n|-----|-----|-----|\n| 内容1 | 内容2 | 内容3 |\n\n`
  insertTextAtCursor(table)
}

const insertTaskList = () => {
  const selectedText = getSelectedText()
  if (selectedText) {
    // 如果有选中文本，转换为任务项
    const task = `- [ ] ${selectedText}\n`
    replaceSelectedText(task)
  } else {
    // 如果没有选中文本，插入模板
    const task = `- [ ] 任务项\n`
    insertTextAtCursor(task)
    selectTextInEditor('任务项')
  }
}

const insertMermaidChart = (chartType) => {
  // 根据图表类型获取对应的Mermaid代码示例
  const chartExamples = {
    flowchart: `graph TD
    A[开始] --> B{条件判断}
    B -->|是| C[执行操作]
    B -->|否| D[结束]
    C --> D`,
    sequence: `sequenceDiagram
    participant 用户
    participant 系统
    
    用户->>系统: 发送请求
    系统->>系统: 处理请求
    系统-->>用户: 返回响应`,
    class: `classDiagram
    class 动物 {
        +名称: string
        +年龄: int
        +吃(): void
        +睡(): void
    }
    
    class 狗 {
        +品种: string
        +叫(): void
    }
    
    狗 --> 动物`,
    state: `stateDiagram
    [*] --> 待机
    待机 --> 运行: 启动
    运行 --> 暂停: 暂停
    暂停 --> 运行: 继续
    运行 --> 停止: 停止
    暂停 --> 停止: 停止
    停止 --> [*]`,
    er: `erDiagram
    用户 ||--o{ 订单 : 下单
    订单 ||--o{ 订单详情 : 包含
    商品 }o--o{ 订单详情 : 组成
    分类 ||--o{ 商品 : 分类`,
    journey: `journey
    title 用户购物旅程
    section 浏览商品
      打开应用: 5: 用户
      搜索商品: 3: 用户
      浏览列表: 7: 用户
    section 下单流程
      查看详情: 4: 用户
      添加购物车: 2: 用户
      结算: 3: 用户
    section 完成购买
      支付: 5: 用户
      确认订单: 2: 用户`,
    gantt: `gantt
    title 项目进度计划
    dateFormat  YYYY-MM-DD
    section 设计
    需求分析: done, des1, 2024-01-01, 3d
    原型设计: des2, after des1, 2d
    设计评审: des3, after des2, 1d
    section 开发
    后端开发: dev1, after des3, 5d
    前端开发: dev2, after des3, 4d
    集成测试: test1, after dev1, 3d`,
    pie: `pie
    title 项目资源分配
    "设计": 30
    "开发": 50
    "测试": 15
    "部署": 5`,
    mindmap: `mindmap
    项目管理
      需求分析
        用户访谈
        竞品分析
        需求文档
      设计阶段
        UI设计
        UX设计
        原型设计
      开发阶段
        后端开发
        前端开发
        测试
      部署上线
        服务器配置
        性能优化
        监控`,
    timeline: `timeline
    title 项目里程碑
    2024-01-01 : 项目启动
    2024-01-10 : 需求确认
    2024-01-20 : 设计完成
    2024-02-10 : 开发完成
    2024-02-20 : 测试通过
    2024-02-25 : 项目上线`
  }
  
  const chartExample = chartExamples[chartType] || chartExamples.flowchart
  
  // 插入Mermaid代码块
  const mermaidCode = `\`\`\`mermaid\n${chartExample}\n\`\`\`\n\n`
  insertTextAtCursor(mermaidCode)
  
  // 选择插入的图表内容
  selectTextInEditor(chartExample)
  
  // 关闭下拉菜单
  showMermaidDropdown.value = false
}

// 保留原有的insertMermaid函数作为默认插入
const insertMermaid = () => {
  insertMermaidChart('flowchart')
}

// 获取选中的文本
const getSelectedText = () => {
  // 尝试使用CodeMirror方式
  const editor = document.querySelector('.bytemd-editor .CodeMirror')
  if (editor && editor.CodeMirror) {
    const cm = editor.CodeMirror
    return cm.getSelection()
  }
  
  // 尝试使用原生textarea方式
  const textarea = document.querySelector('.bytemd-editor .markdown-textarea')
  if (textarea && textarea.selectionStart !== textarea.selectionEnd) {
    return textarea.value.substring(textarea.selectionStart, textarea.selectionEnd)
  }
  
  return ''
}

// 替换选中的文本
const replaceSelectedText = (text) => {
  // 尝试使用CodeMirror方式
  const editor = document.querySelector('.bytemd-editor .CodeMirror')
  if (editor && editor.CodeMirror) {
    const cm = editor.CodeMirror
    const selection = cm.getSelection()
    if (selection) {
      cm.replaceSelection(text)
      cm.focus()
      return true
    }
  }
  
  // 尝试使用原生textarea方式
  const textarea = document.querySelector('.bytemd-editor .markdown-textarea')
  if (textarea) {
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    if (start !== end) {
      const newValue = textarea.value.substring(0, start) + text + textarea.value.substring(end)
      textarea.value = newValue
      textarea.focus()
      content.value = newValue
      emit('update:modelValue', newValue)
      return true
    }
  }
  
  return false
}

// 选中编辑器中的文本
const selectTextInEditor = (text) => {
  // 尝试使用CodeMirror方式
  const editor = document.querySelector('.bytemd-editor .CodeMirror')
  if (editor && editor.CodeMirror) {
    const cm = editor.CodeMirror
    const cursor = cm.getCursor()
    
    // 计算文本的行数和字符数
    const lines = text.split('\n')
    const textLength = text.length
    const linesCount = lines.length
    
    // 计算起始位置
    let startLine = cursor.line
    let startCh = cursor.ch - textLength
    
    // 如果起始位置在当前行之前，需要调整
    if (startCh < 0) {
      startLine = Math.max(0, cursor.line - linesCount + 1)
      startCh = Math.max(0, cursor.ch - textLength + (cursor.line - startLine) * lines[0].length)
    }
    
    const start = { line: startLine, ch: startCh }
    const end = cursor
    
    cm.setSelection(start, end)
    cm.focus()
    return true
  }
  
  // 尝试使用原生textarea方式
  const textarea = document.querySelector('.bytemd-editor .markdown-textarea')
  if (textarea) {
    const cursorPos = textarea.selectionStart
    const startPos = cursorPos - text.length
    if (startPos >= 0) {
      textarea.setSelectionRange(startPos, cursorPos)
      textarea.focus()
      return true
    }
  }
  
  return false
}

// 处理拖拽事件
const handleDragOver = (event) => {
  event.preventDefault()
  event.dataTransfer.dropEffect = 'copy'
}

const handleDragEnter = (event) => {
  event.preventDefault()
  isDragOver.value = true
}

const handleDragLeave = (event) => {
  // 只有当离开整个编辑器区域时才隐藏拖拽状态
  if (!event.currentTarget.contains(event.relatedTarget)) {
    isDragOver.value = false
  }
}

// 分隔线样式已在CSS中添加，无需额外JavaScript功能

// 处理拖拽图片上传
const handleImageDrop = async (event) => {
  // 处理图片拖拽放下事件
  isDragOver.value = false
  
  const files = event.dataTransfer.files
  
  if (files.length === 0) {
    return
  }
  
  const file = files[0]
  // 验证文件类型并处理上传
  
  // 验证文件类型
  if (!file.type.startsWith('image/')) {
    alert('请拖拽图片文件')
    return
  }
  
  // 验证文件大小 (限制为10MB)
  if (file.size > 10 * 1024 * 1024) {
    alert('图片文件过大，请选择小于10MB的图片')
    return
  }
  
  // 使用通用上传函数处理图片上传
  await uploadImage(file)
}

// 在光标位置插入文本的辅助函数
const insertTextAtCursor = (text) => {
  // 尝试使用CodeMirror方式
  const editor = document.querySelector('.bytemd-editor .CodeMirror')
  if (editor && editor.CodeMirror) {
    const cm = editor.CodeMirror
    const cursor = cm.getCursor()
    const doc = cm.getDoc()
    
    // 在光标位置插入文本
    doc.replaceRange(text, cursor)
    
    // 移动光标到插入文本的末尾
    const newCursor = {
      line: cursor.line,
      ch: cursor.ch + text.length
    }
    cm.setCursor(newCursor)
    cm.focus()
    return true
  }
  
  // 尝试使用原生textarea方式
  const textarea = document.querySelector('.bytemd-editor .markdown-textarea')
  if (textarea) {
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    const newValue = textarea.value.substring(0, start) + text + textarea.value.substring(end)
    
    textarea.value = newValue
    content.value = newValue
    emit('update:modelValue', newValue)
    
    // 移动光标到插入文本的末尾
    const newPosition = start + text.length
    textarea.setSelectionRange(newPosition, newPosition)
    textarea.focus()
    return true
  }
  
  // 如果都不行，直接追加到内容末尾
  content.value += text
  emit('update:modelValue', content.value)
  return false
}

// 监听props变化
watch(() => props.title, (newTitle) => {
  title.value = newTitle
})

watch(() => props.modelValue, (newContent) => {
  content.value = newContent || ''
  generateToc()
})

// 初始化
onMounted(() => {
  generateToc()
  
  // 如果同步滚动已经开启且是双屏模式，设置同步滚动
  if (syncScroll.value && editorMode.value === 'split') {
    setupSyncScroll()
  }
  
  // 添加全局拖拽事件监听器
  document.addEventListener('dragover', (e) => {
    e.preventDefault()
  })
  
  document.addEventListener('drop', (e) => {
    e.preventDefault()
  })
})

// 组件卸载时清理
onUnmounted(() => {
  if (syncScrollCleanup) {
    syncScrollCleanup()
    syncScrollCleanup = null
  }
})
</script>

<style scoped>
/* Mermaid下拉菜单样式 */
.mermaid-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 0;
  padding-top: 4px;
  padding-bottom: 2px;
  background: white;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  min-width: 100px;
  z-index: 1000;
  font-size: 12px;
}

.mermaid-option {
  padding: 4px 10px;
  cursor: pointer;
  white-space: nowrap;
  line-height: 1.4;
}

.mermaid-option:hover {
  background-color: #f5f5f5;
}

/* 掘金编辑器整体样式 */
.juejin-editor {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: #fff;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  overflow: hidden;
}

/* 顶部标题栏 */
.juejin-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 24px;
  border-bottom: 1px solid #e9ecef;
  background: #fff;
  z-index: 100;
}

.header-left {
  flex: 1;
  margin-right: 24px;
}

.title-input {
  width: 100%;
  border: none;
  outline: none;
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  background: transparent;
}

.title-input::placeholder {
  color: #86909c;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.status-text {
  color: #86909c;
  font-size: 14px;
}

.draft-btn, .publish-btn {
  padding: 8px 16px;
  border: 1px solid #e9ecef;
  border-radius: 6px;
  background: #fff;
  color: #1a1a1a;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.publish-btn {
  background: #007bff;
  color: #fff;
  border-color: #007bff;
}

.draft-btn:hover, .publish-btn:hover {
  opacity: 0.8;
}

.editor-switcher {
  position: relative;
}

.toggle-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #e9ecef;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
}

.toggle-btn:hover {
  background: #f8f9fa;
}

.user-navigator {
  position: relative;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #007bff;
  cursor: pointer;
}

/* 主编辑器区域 */
.juejin-main {
  flex: 1;
  overflow: hidden;
  min-height: 0;
  display: flex;
  flex-direction: column;
}

.bytemd {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.bytemd-body {
  flex: 1;
  display: flex;
  overflow: hidden;
  min-height: 0;
  width: 100%;
}



/* 工具栏 */
.bytemd-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 4px 12px;
  border-bottom: 1px solid #e9ecef;
  background: #f8f9fa;
}

.bytemd-toolbar-left {
  display: flex;
  align-items: center;
  gap: 4px;
}

.bytemd-toolbar-right {
  display: flex;
  align-items: center;
  gap: 4px;
}

.bytemd-toolbar-icon {
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  cursor: pointer;
  color: #666;
  transition: all 0.2s;
}

/* 有下拉菜单的工具栏图标特殊样式 */
.bytemd-toolbar-item .bytemd-toolbar-icon:hover {
  background-color: #f0f0f0;
  color: #007bff;
}

.bytemd-toolbar-item {
  position: relative;
}

.relative {
  position: relative;
}

/* 标题级别下拉菜单样式 */
.heading-dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  margin-top: 0;
  padding-top: 2px;
  padding-bottom: 2px;
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  z-index: 1000;
  min-width: 120px;
  overflow: hidden;
  transform-origin: top left;
  animation: dropdownFadeIn 0.2s ease-out;
}

/* 添加伪元素创建桥接区域，确保鼠标可以平滑移动到下拉菜单 */
.bytemd-toolbar-item.relative::before {
  content: '';
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  height: 8px;
  z-index: 999;
  background: transparent;
}

/* 下拉菜单淡入动画 */
@keyframes dropdownFadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

.heading-option {
  display: flex;
  align-items: center;
  padding: 4px 12px;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 12px;
  position: relative;
}

.heading-option:hover {
  background-color: #f8f9fa;
  padding-left: 16px;
  padding-right: 8px;
}

/* 选中效果 */
.heading-option:active {
  background-color: #e9ecef;
}

/* 选中时的左侧指示器 */
.heading-option:hover::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 3px;
  background-color: #007bff;
}

.heading-level {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 20px;
  background-color: #f0f0f0;
  border-radius: 4px;
  margin-right: 8px;
  font-weight: bold;
  color: #333;
}

.heading-preview {
  flex: 1;
  color: #333;
}

/* 不同级别的标题预览样式 */
.heading-option:nth-child(1) .heading-preview { font-size: 18px; font-weight: bold; }
.heading-option:nth-child(2) .heading-preview { font-size: 16px; font-weight: bold; }
.heading-option:nth-child(3) .heading-preview { font-size: 14px; font-weight: bold; }
.heading-option:nth-child(4) .heading-preview { font-size: 13px; font-weight: bold; }
.heading-option:nth-child(5) .heading-preview { font-size: 12px; font-weight: bold; }
.heading-option:nth-child(6) .heading-preview { font-size: 11px; font-weight: bold; }

.bytemd-toolbar-icon:hover {
  background: #e9ecef;
  color: #1a1a1a;
}

.bytemd-toolbar-icon-active {
  background: #007bff;
  color: #fff;
}



/* 分屏模式容器 */
.bytemd-split {
  display: flex;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

/* 编辑区容器 - 作为滚动容器 */
.bytemd-editor {
  border-right: 1px solid #e9ecef;
  width: 100%;
  height: 100%;
  background: #fafafa;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
}

/* 预览区容器 - 作为滚动容器 */
.bytemd-preview {
  width: 100%;
  height: 100%;
  background: #fff;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
}

/* 编辑区内容器 - 用于内容居中 */
.editor-content-wrapper {
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  height: 100%;
}

/* 预览区内容器 - 用于内容居中 */
.preview-content-wrapper {
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  height: 100%;
}

/* 编辑区文本域 */
.markdown-textarea {
  max-width: 800px;
  width: 100%;
  min-height: calc(100vh - 200px);
  border: none;
  outline: none;
  padding: 24px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 14px;
  line-height: 1.6;
  color: #1a1a1a;
  background: transparent;
  resize: none;
  box-sizing: border-box;
  margin: 0;
}

/* CodeMirror编辑器样式 */
.CodeMirror {
  height: 100% !important;
  width: 100% !important;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace !important;
  font-size: 14px !important;
  line-height: 1.6 !important;
}

.CodeMirror-scroll {
  height: 100% !important;
  overflow: hidden !important; /* 隐藏CodeMirror的滚动条 */
}

.CodeMirror-sizer {
  min-height: calc(100vh - 200px) !important;
}

.markdown-textarea::placeholder {
  color: #86909c;
}

/* 预览区域样式 - 使用专业的GitHub markdown主题 */
.markdown-body {
  max-width: 800px;
  width: 100%;
  min-height: calc(100vh - 200px);
  padding: 24px;
  box-sizing: border-box;
  /* GitHub markdown主题已经包含了所有必要的样式 */
}

/* 确保Mermaid图表在GitHub主题中正确显示的额外样式 */
.markdown-body .mermaid {
  margin: 24px 0;
  overflow-x: auto;
  text-align: center;
}

.markdown-body .mermaid-rendered {
  text-align: center;
  margin: 24px 0;
  padding: 16px;
  background: #f6f8fa;
  border-radius: 6px;
}

/* 确保任务列表插件与GitHub主题兼容 */
.markdown-body .task-list-item {
  list-style-type: none;
}

.markdown-body .task-list-item input[type="checkbox"] {
  margin-right: 8px;
}

.markdown-body code {
  padding: 0.2em 0.4em;
  margin: 0;
  font-size: 85%;
  background-color: rgba(27,31,35,0.05);
  border-radius: 3px;
}

/* 侧边栏 */
.bytemd-sidebar {
  width: 280px;
  border-left: 1px solid #e9ecef;
  background: #fff;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.bytemd-sidebar-close {
  padding: 16px;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: flex-end;
  cursor: pointer;
}

.bytemd-help {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}

.bytemd-help h2 {
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 16px 0;
  color: #1a1a1a;
}

.bytemd-help ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.bytemd-help li {
  display: flex;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f0f0f0;
}

.bytemd-help-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
  color: #86909c;
}

.bytemd-help-title {
  flex: 1;
  font-size: 14px;
  color: #1a1a1a;
}

.bytemd-help-content {
  font-size: 12px;
  color: #86909c;
}

.bytemd-help-content code {
  background: #f0f0f0;
  padding: 2px 6px;
  border-radius: 3px;
  font-family: monospace;
}

/* 拖拽上传样式 */
.bytemd-editor.drag-over {
  position: relative;
  border: 1px solid #333 !important;
  background: rgba(255, 255, 255, 0.8) !important;
  transition: all 0.3s ease;
}

.drag-overlay {
  position: absolute;
  top: 20%;
  left: 20%;
  right: 20%;
  bottom: 20%;
  background: rgba(255, 255, 255, 0.8);
  border: 1px solid #333;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  pointer-events: none;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

.drag-overlay-content {
  text-align: center;
  color: #333;
  padding: 20px;
  max-width: 300px;
  width: 90%;
}

.drag-overlay-content svg {
  margin-bottom: 12px;
  width: 36px;
  height: 36px;
  stroke-width: 1.5;
  fill: none;
  stroke: #333;
}

.drag-overlay-content p {
  margin: 0;
  font-size: 14px;
  font-weight: 400;
  margin-bottom: 12px;
  color: #333;
  line-height: 1.4;
}

.select-image-btn {
  background: linear-gradient(135deg, #1677ff 0%, #0958d9 100%);
  color: white;
  border: none;
  padding: 12px 28px;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(22, 119, 255, 0.3);
  pointer-events: all;
  text-transform: none;
}

.select-image-btn:hover {
  background: linear-gradient(135deg, #4096ff 0%, #1677ff 100%);
  box-shadow: 0 6px 16px rgba(22, 119, 255, 0.4);
  transform: translateY(-1px);
}

.select-image-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 8px rgba(22, 119, 255, 0.3);
}

/* 关闭按钮样式 */
.drag-overlay-close {
    position: absolute;
    top: 12px;
    right: 12px;
    width: 32px;
    height: 32px;
    border-radius: 50%;
    background: rgba(0, 0, 0, 0.06);
    border: none;
    color: #666;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    padding: 0;
    pointer-events: all;
  }
  
  .drag-overlay-close:hover {
    background: rgba(0, 0, 0, 0.1);
    color: #333;
    transform: scale(1.1);
  }
  
  .drag-overlay-close:active {
    transform: scale(0.95);
  }
  
  /* 提示文字样式 */
  .drag-overlay-hint {
    margin-top: 12px;
    font-size: 12px;
    color: #8c8c8c;
    font-weight: 400;
  }
  
  /* 拖拽时的动画效果 */
  .bytemd-editor.drag-over .drag-overlay-content .upload-icon {
    animation: bounce 1s infinite;
  }
  
  @keyframes bounce {
    0%, 100% {
      transform: translateY(0);
    }
    50% {
      transform: translateY(-10px);
    }
  }

.bytemd-help-content kbd {
  background: #f0f0f0;
  padding: 2px 6px;
  border-radius: 3px;
  font-family: monospace;
  border: 1px solid #d0d0d0;
}

.bytemd-toc {
  padding: 16px;
  border-top: 1px solid #e9ecef;
}

.bytemd-toc h2 {
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 16px 0;
  color: #1a1a1a;
}

.bytemd-toc ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.bytemd-toc li {
  margin-bottom: 4px;
}

.bytemd-toc a {
  display: block;
  padding: 4px 8px;
  color: #86909c;
  text-decoration: none;
  border-radius: 4px;
  transition: all 0.2s;
  cursor: pointer;
}

.bytemd-toc a:hover {
  background: #f0f0f0;
  color: #1a1a1a;
}

.toc-level-1 { padding-left: 0; }
.toc-level-2 { padding-left: 16px; }
.toc-level-3 { padding-left: 32px; }
.toc-level-4 { padding-left: 48px; }
.toc-level-5 { padding-left: 64px; }
.toc-level-6 { padding-left: 80px; }

/* 状态栏样式 */
.bytemd-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 16px;
  background-color: #f8f9fa;
  border-top: 1px solid #e9ecef;
  font-size: 12px;
  color: #666;
  min-height: 32px;
  flex-shrink: 0;
}

.bytemd-status-left {
  display: flex;
  gap: 16px;
}

.bytemd-status-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.bytemd-status-right label {
  display: flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
}

.bytemd-status-right input[type="checkbox"] {
  margin: 0;
}

.bytemd-status-right span {
  cursor: pointer;
  color: #007bff;
  transition: color 0.2s;
}

.bytemd-status-right span:hover {
  color: #0056b3;
}



/* 响应式设计 */
@media (max-width: 768px) {
  .juejin-header {
    padding: 12px 16px;
  }
  
  .title-input {
    font-size: 20px;
  }
  
  .bytemd-toolbar {
    padding: 8px 12px;
  }
  
  .bytemd-toolbar-icon {
    width: 28px;
    height: 28px;
  }
  
  .markdown-textarea {
    padding: 16px;
  }
  
  .markdown-body {
    padding: 16px;
  }
  
  /* Mermaid下拉菜单样式 */
  .mermaid-dropdown {
    font-size: 11px !important;
    min-width: 90px !important;
  }
  
  .mermaid-option {
    padding: 3px 10px !important;
    line-height: 1.3 !important;
    white-space: nowrap;
  }
  
  /* 标题下拉菜单样式 */
  .heading-dropdown {
    font-size: 12px;
  }
  
  .heading-option {
    padding: 4px 12px;
    line-height: 1.4;
  }
}
</style>
