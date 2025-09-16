<template>
  <div class="admin-panel">
    <h1>管理员面板</h1>
    
    <div v-if="loading" class="loading">加载中...</div>
    
    <div v-else-if="error" class="error-message">{{ error }}</div>
    
    <div v-else-if="users.length === 0" class="empty">暂无用户</div>
    
    <div v-else class="users-table-container">
      <table class="users-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>用户名</th>
            <th>邮箱</th>
            <th>角色</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.username }}</td>
            <td>{{ user.email || '未设置' }}</td>
            <td>
              <select v-model="user.editingRole" @change="updateUserRole(user.id, user.editingRole)" :disabled="editingUser === user.id">
                <option value="user">普通用户</option>
                <option value="admin">管理员</option>
              </select>
            </td>
            <td>{{ formatDate(user.createdAt) }}</td>
            <td>
              <button 
                class="edit-button"
                @click="toggleEditUser(user.id)"
                :disabled="editingUser === user.id"
              >
                {{ editingUser === user.id ? '保存' : '编辑' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import blogApi from '../services/api.js'

export default {
  name: 'AdminPanel',
  data() {
    return {
      users: [],
      loading: false,
      error: '',
      editingUser: null,
      originalUserData: {}
    }
  },
  mounted() {
    this.fetchUsers()
  },
  methods: {
    async fetchUsers() {
      this.loading = true
      this.error = ''
      
      try {
        const response = await blogApi.getAllUsers()
        // 为每个用户添加编辑状态的角色
        this.users = response.data.map(user => ({
          ...user,
          editingRole: user.role
        }))
      } catch (error) {
        this.error = error.response?.data?.error || '获取用户列表失败'
        console.error('获取用户列表失败:', error)
      } finally {
        this.loading = false
      }
    },
    
    toggleEditUser(userId) {
      const user = this.users.find(u => u.id === userId)
      if (!user) return
      
      if (this.editingUser === userId) {
        // 取消编辑，恢复原始数据
        this.cancelEditUser(userId)
      } else {
        // 开始编辑，保存原始数据
        this.originalUserData[userId] = {
          role: user.role
        }
        this.editingUser = userId
      }
    },
    
    cancelEditUser(userId) {
      const user = this.users.find(u => u.id === userId)
      if (user && this.originalUserData[userId]) {
        user.editingRole = this.originalUserData[userId].role
        delete this.originalUserData[userId]
      }
      this.editingUser = null
    },
    
    async updateUserRole(userId, newRole) {
      this.loading = true
      this.error = ''
      
      try {
        await blogApi.updateUser(userId, { role: newRole })
        
        // 更新本地数据
        const user = this.users.find(u => u.id === userId)
        if (user) {
          user.role = newRole
          user.editingRole = newRole
        }
        
        this.editingUser = null
        delete this.originalUserData[userId]
        
        // 显示成功提示
        alert('用户角色更新成功')
      } catch (error) {
        this.error = error.response?.data?.error || '更新用户角色失败'
        console.error('更新用户角色失败:', error)
        // 恢复原始角色
        this.cancelEditUser(userId)
      } finally {
        this.loading = false
      }
    },
    
    formatDate(dateString) {
      const date = new Date(dateString)
      return date.toLocaleString('zh-CN')
    }
  }
}
</script>

<style scoped>
.admin-panel {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.admin-panel h1 {
  text-align: center;
  margin-bottom: 30px;
  color: #333;
}

.loading {
  text-align: center;
  padding: 40px;
  color: #666;
}

.error-message {
  background-color: #f8d7da;
  color: #721c24;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #f5c6cb;
  text-align: center;
}

.empty {
  text-align: center;
  padding: 40px;
  color: #666;
}

.users-table-container {
  overflow-x: auto;
}

.users-table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

.users-table th,
.users-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.users-table th {
  background-color: #f8f9fa;
  font-weight: 600;
  color: #495057;
}

.users-table tr:hover {
  background-color: #f8f9fa;
}

.users-table select {
  padding: 6px 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  background-color: white;
}

.edit-button {
  padding: 6px 12px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s;
}

.edit-button:hover:not(:disabled) {
  background-color: #218838;
}

.edit-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

@media (max-width: 768px) {
  .users-table {
    font-size: 14px;
  }
  
  .users-table th,
  .users-table td {
    padding: 8px 10px;
  }
}
</style>