/* 页面结构 */
<template>
  <div class="app"> <!-- 页面根容器 -->
    <header class="app-header">
      <h1><router-link to="/">我的博客系统</router-link></h1>
      <!-- 解释：
        - <h1>是大标题“我的博客系统”
        - <router-link to="/"> 是Vue里的“链接工具”，类似网页里的<a href="/">，但点击后不会刷新整个页面
        - to="/" 表示点击后跳转到网站的首页
      -->
    </header>
    <main class="app-main">
      <router-view></router-view>
      <!-- 解释：
        - 这是Vue里的“内容展示区”，相当于一个“动态窗口”
        - 当你点击上面的链接（比如首页、文章页），这里会自动显示对应页面的内容
        - 比如：点击“首页”，这里显示博客列表；点击“某篇文章”，这里显示文章详情
      -->
    </main>
    <footer class="app-footer">
      <p>&copy; 2024 我的博客系统. 保留所有权利.</p>
    </footer>
  </div>
</template>

/* 全局样式 */
<style>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  line-height: 1.6;
  color: #333;
  background-color: #f5f5f5;
}

.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  background-color: #333;
  color: white;
  padding: 15px 0;
  text-align: center;
}

.app-header h1 {
  margin: 0;
  font-size: 24px;
}

.app-header a {
  color: white;
  text-decoration: none;
}

.app-main {
  flex: 1;
  padding: 20px;
}

.app-footer {
  background-color: #333;
  color: white;
  text-align: center;
  padding: 15px 0;
  font-size: 14px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .app-header h1 {
    font-size: 20px;
  }
  
  .app-main {
    padding: 10px;
  }
}
</style>
