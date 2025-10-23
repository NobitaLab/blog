/* 页面结构 */
<template>
  <div class="app"> <!-- 页面根容器 -->
    <header class="app-header">
      <nav class="nav-container">
        <div class="logo-container">
          <router-link to="/" class="logo">我的博客系统</router-link>
        </div>
        
        <div class="nav-links">
          <!-- 公开导航链接 -->
          <router-link to="/" class="nav-link">首页</router-link>
          
          <!-- 登录用户可见的导航链接 -->
          <router-link v-if="isAuthenticated" to="/create" class="nav-link">写博客</router-link>
          
          <!-- 管理员可见的导航链接 -->
          <router-link v-if="isAdmin" to="/admin" class="nav-link admin-link">管理员面板</router-link>
          
          <!-- 用户信息和操作 -->
          <div v-if="isAuthenticated" class="user-section">
            <span class="welcome">欢迎, {{ username }}</span>
            <span class="role-badge" v-if="isAdmin">管理员</span>
            <button class="logout-button" @click="handleLogout">退出登录</button>
          </div>
          
          <!-- 未登录用户可见的导航链接 -->
          <div v-else class="auth-links">
            <router-link to="/login" class="nav-link">登录</router-link>
            <router-link to="/register" class="nav-link">注册</router-link>
          </div>
        </div>
      </nav>
    </header>
    <main class="app-main">
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
import blogApi from './services/api.js'

export default {
  name: 'App',
  data() {
    return {
      isAuthenticated: false,
      username: '',
      isAdmin: false
    }
  },
  mounted() {
    this.checkUserStatus()
    // 监听路由变化，更新用户状态
    this.$router.beforeEach((to, from, next) => {
      this.checkUserStatus()
      next()
    })
  },
  methods: {
    checkUserStatus() {
      const token = localStorage.getItem('token')
      const userInfo = localStorage.getItem('userInfo')
      
      if (token && userInfo) {
        try {
          const user = JSON.parse(userInfo)
          this.isAuthenticated = true
          this.username = user.username
          this.isAdmin = user.role === 'admin'
        } catch (error) {
          this.logout()
        }
      } else {
        this.logout()
      }
    },
    
    async handleLogout() {
      try {
        await blogApi.logout()
      } catch (error) {
        console.error('注销失败:', error)
      } finally {
        this.logout()
        this.$router.push('/')
      }
    },
    
    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      this.isAuthenticated = false
      this.username = ''
      this.isAdmin = false
    }
  }
}
</script>

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
  padding: 0;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 50px;
}

.logo-container {
  flex-shrink: 0;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: white;
  text-decoration: none;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 20px;
  flex-wrap: wrap;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 8px 15px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-link:hover {
  background-color: #555;
}

.admin-link {
  color: #ffd700;
  font-weight: 500;
}

.user-section {
  display: flex;
  align-items: center;
  gap: 15px;
}

.welcome {
  color: white;
  font-weight: 500;
}

.role-badge {
  background-color: #ffd700;
  color: #333;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: bold;
}

.logout-button {
  background-color: #dc3545;
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.3s;
}

.logout-button:hover {
  background-color: #c82333;
}

.auth-links {
  display: flex;
  gap: 15px;
}

.app-main {
  flex: 1;
  padding: 20px;
}


/* 响应式设计 */
@media (max-width: 768px) {
  .nav-container {
    flex-direction: column;
    height: auto;
    padding: 15px 20px;
  }
  
  .nav-links {
    margin-top: 15px;
    justify-content: center;
    flex-direction: column;
    width: 100%;
    gap: 10px;
  }
  
  .user-section,
  .auth-links {
    flex-direction: column;
    width: 100%;
    gap: 10px;
    text-align: center;
  }
}
</style>
