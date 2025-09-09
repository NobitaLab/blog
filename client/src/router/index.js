import { createRouter, createWebHistory } from 'vue-router'
import BlogList from '../views/BlogList.vue'
import BlogDetail from '../views/BlogDetail.vue'
import BlogEdit from '../views/BlogEdit.vue'
import BlogCreate from '../views/BlogCreate.vue'

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
    props: true
  },
  {
    path: '/create',
    name: 'BlogCreate',
    component: BlogCreate
  }
]

const router = createRouter({
  history: createWebHistory('/'),
  routes
})

export default router