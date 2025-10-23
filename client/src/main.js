import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
// 导入专业的GitHub markdown CSS主题
import 'github-markdown-css/github-markdown.css'
// 导入代码高亮样式
import 'highlight.js/styles/github.css'

const app = createApp(App)
app.use(router)
app.mount('#app')

export { router }
