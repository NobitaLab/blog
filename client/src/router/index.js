import { createRouter, createWebHistory } from 'vue-router'
import BlogList from '../views/BlogList.vue'
import BlogDetail from '../views/BlogDetail.vue'
import BlogEdit from '../views/BlogEdit.vue'
import BlogCreate from '../views/BlogCreate.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import AdminPanel from '../views/AdminPanel.vue'
import blogApi from '../services/api.js'

const routes = [
  {
    path: '/',
    name: 'BlogList',
    component: BlogList
  },
  {
    path: '/blog/:id',
    name: 'BlogDetail',
    component: BlogDetail,
    props: true
  },
  {
    path: '/edit/:id',
    name: 'BlogEdit',
    component: BlogEdit,
    props: true,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/create',
    name: 'BlogCreate',
    component: BlogCreate,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      guestOnly: true
    }
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: {
      guestOnly: true
    }
  },
  {
    path: '/admin',
    name: 'AdminPanel',
    component: AdminPanel,
    meta: {
      requiresAuth: true,
      requiresAdmin: true
    }
  },
]

const router = createRouter({
  history: createWebHistory('/'),
  routes
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  // 优先检查localStorage中的token和用户信息，避免不必要的API调用
  const token = localStorage.getItem('token')
  let isAuthenticated = !!token
  let userRole = 'user'
  
  // 只有在需要认证的路由或已登录状态下，才调用API验证用户状态
  if ((to.meta.requiresAuth || to.meta.requiresAdmin) && token) {
    try {
      const response = await blogApi.getCurrentUser()
      if (response.data) {
        isAuthenticated = true
        userRole = response.data.role
        // 保存用户信息到localStorage
        localStorage.setItem('userInfo', JSON.stringify(response.data))
      } else {
        // API返回数据为空，视为认证失败
        isAuthenticated = false
        localStorage.removeItem('token')
        localStorage.removeItem('userInfo')
      }
    } catch (error) {
      // API调用失败，视为认证失败
      isAuthenticated = false
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
    }
  }
  
  // 如果路由需要认证但用户未登录，跳转到登录页
  if (to.meta.requiresAuth && !isAuthenticated) {
    return next('/login')
  }
  
  // 如果路由只对访客开放但用户已登录，跳转到首页
  if (to.meta.guestOnly && isAuthenticated) {
    return next('/')
  }
  
  // 如果路由需要管理员权限但用户不是管理员，显示权限不足
  if (to.meta.requiresAdmin && userRole !== 'admin') {
    alert('需要管理员权限访问此页面')
    return next(from.path || '/')
  }
  
  next()
})

export default router