<template>
  <div class="collaboration-editor">
    <div class="editor-header">
      <h3>{{ noteTitle }}</h3>
      <div class="active-users">
        <span>在线用户: {{ activeUsers.length }}</span>
        <div class="user-list">
          <span v-for="user in activeUsers" :key="user.id" class="user-item">
            {{ user.username }}
          </span>
        </div>
      </div>
    </div>
    
    <div class="editor-container">
      <textarea
        v-model="content"
        @input="handleInput"
        @keydown="handleKeydown"
        @keyup="handleKeyup"
        placeholder="开始协同编辑..."
        class="editor-textarea"
      ></textarea>
    </div>
    
    <div class="editor-footer">
      <div class="like-section">
        <button @click="toggleLike" :class="{ 'liked': isLiked }">
          {{ isLiked ? '取消点赞' : '点赞' }} ({{ likeCount }})
        </button>
      </div>
      <div class="status">
        <span v-if="isConnected" class="connected">已连接</span>
        <span v-else class="disconnected">未连接</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CollaborationEditor',
  props: {
    noteId: {
      type: Number,
      required: true
    },
    noteTitle: {
      type: String,
      default: '协同编辑'
    },
    initialContent: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      content: this.initialContent,
      ws: null,
      isConnected: false,
      activeUsers: [],
      operations: [],
      isLiked: false,
      likeCount: 0,
      cursorPosition: 0,
      lastContent: this.initialContent
    }
  },
  mounted() {
    this.initWebSocket()
    this.loadLikeStatus()
  },
  beforeUnmount() {
    this.closeWebSocket()
  },
  methods: {
    initWebSocket() {
      const userInfo = this.getUserInfo()
      const wsUrl = `ws://localhost:8888/api/v1/collaboration/ws?noteId=${this.noteId}&userId=${userInfo.id}&username=${userInfo.username}`
      
      this.ws = new WebSocket(wsUrl)
      
      this.ws.onopen = () => {
        this.isConnected = true
        console.log('WebSocket连接已建立')
      }
      
      this.ws.onmessage = (event) => {
        const data = JSON.parse(event.data)
        this.handleWebSocketMessage(data)
      }
      
      this.ws.onclose = () => {
        this.isConnected = false
        console.log('WebSocket连接已关闭')
        // 尝试重连
        setTimeout(() => {
          this.initWebSocket()
        }, 3000)
      }
      
      this.ws.onerror = (error) => {
        console.error('WebSocket错误:', error)
        this.isConnected = false
      }
    },
    
    handleWebSocketMessage(data) {
      switch (data.type) {
        case 'history':
          this.operations = data.operations || []
          this.applyOperations()
          break
        case 'active_users':
          this.activeUsers = data.activeUsers || []
          break
        case 'user_joined':
          this.addActiveUser(data)
          break
        case 'user_left':
          this.removeActiveUser(data.sessionId)
          break
        case 'operation':
          this.handleRemoteOperation(data.operation)
          break
        case 'pong':
          // 心跳响应
          break
      }
    },
    
    handleInput(event) {
      const newContent = event.target.value
      const position = event.target.selectionStart
      
      // 检测变化类型
      if (newContent.length > this.lastContent.length) {
        // 插入操作
        const insertedText = newContent.substring(this.lastContent.length)
        this.sendOperation('insert', position - insertedText.length, insertedText)
      } else if (newContent.length < this.lastContent.length) {
        // 删除操作
        const deletedText = this.lastContent.substring(0, this.lastContent.length - newContent.length)
        this.sendOperation('delete', position, deletedText)
      }
      
      this.lastContent = newContent
      this.cursorPosition = position
    },
    
    handleKeydown(event) {
      // 记录按键位置
      this.cursorPosition = event.target.selectionStart
    },
    
    handleKeyup(event) {
      // 更新光标位置
      this.cursorPosition = event.target.selectionStart
    },
    
    sendOperation(type, position, content) {
      if (this.ws && this.isConnected) {
        this.ws.send(JSON.stringify({
          type: 'operation',
          opType: type,
          position: position,
          content: content
        }))
      }
    },
    
    handleRemoteOperation(operation) {
      // 应用远程操作
      this.operations.push(operation)
      this.applyOperations()
    },
    
    applyOperations() {
      // 按时间戳排序操作
      this.operations.sort((a, b) => {
        if (a.timestamp !== b.timestamp) {
          return a.timestamp - b.timestamp
        }
        return a.userId - b.userId
      })
      
      // 应用操作到内容
      let result = this.initialContent
      let offset = 0
      
      for (const op of this.operations) {
        switch (op.type) {
          case 'insert':
            const insertPos = op.position + offset
            if (insertPos >= 0 && insertPos <= result.length) {
              result = result.substring(0, insertPos) + op.content + result.substring(insertPos)
              offset += op.content.length
            }
            break
          case 'delete':
            const deletePos = op.position + offset
            if (deletePos >= 0 && deletePos < result.length) {
              const deleteLen = op.content.length
              if (deletePos + deleteLen <= result.length) {
                result = result.substring(0, deletePos) + result.substring(deletePos + deleteLen)
                offset -= deleteLen
              }
            }
            break
        }
      }
      
      this.content = result
      this.lastContent = result
    },
    
    addActiveUser(data) {
      this.activeUsers.push({
        id: data.sessionId,
        username: data.username,
        userId: data.userId
      })
    },
    
    removeActiveUser(sessionId) {
      this.activeUsers = this.activeUsers.filter(user => user.id !== sessionId)
    },
    
    closeWebSocket() {
      if (this.ws) {
        this.ws.close()
      }
    },
    
    async loadLikeStatus() {
      try {
        const [likeStatus, likeCount] = await Promise.all([
          this.checkLikeStatus(),
          this.getLikeCount()
        ])
        this.isLiked = likeStatus
        this.likeCount = likeCount
      } catch (error) {
        console.error('加载点赞状态失败:', error)
      }
    },
    
    async checkLikeStatus() {
      const response = await fetch(`/api/v1/like/check/${this.noteId}`, {
        headers: {
          'Authorization': `Bearer ${this.getToken()}`
        }
      })
      const data = await response.json()
      return data.data.isLiked
    },
    
    async getLikeCount() {
      const response = await fetch(`/api/v1/like/count/${this.noteId}`)
      const data = await response.json()
      return data.data.likeCount
    },
    
    async toggleLike() {
      try {
        const url = `/api/v1/like/${this.isLiked ? 'note' : 'note'}/${this.noteId}`
        const method = this.isLiked ? 'DELETE' : 'POST'
        
        const response = await fetch(url, {
          method: method,
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })
        
        if (response.ok) {
          this.isLiked = !this.isLiked
          this.likeCount += this.isLiked ? 1 : -1
        }
      } catch (error) {
        console.error('点赞操作失败:', error)
      }
    },
    
    getUserInfo() {
      // 从localStorage或vuex获取用户信息
      return {
        id: 1, // 示例用户ID
        username: '用户' + Math.floor(Math.random() * 1000)
      }
    },
    
    getToken() {
      // 从localStorage获取token
      return localStorage.getItem('token') || ''
    }
  }
}
</script>

<style scoped>
.collaboration-editor {
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
}

.editor-header {
  background: #f5f5f5;
  padding: 12px 16px;
  border-bottom: 1px solid #ddd;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.editor-header h3 {
  margin: 0;
  color: #333;
}

.active-users {
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-list {
  display: flex;
  gap: 4px;
}

.user-item {
  background: #007bff;
  color: white;
  padding: 2px 6px;
  border-radius: 12px;
  font-size: 12px;
}

.editor-container {
  padding: 16px;
}

.editor-textarea {
  width: 100%;
  min-height: 300px;
  border: none;
  outline: none;
  resize: vertical;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  line-height: 1.5;
}

.editor-footer {
  background: #f5f5f5;
  padding: 12px 16px;
  border-top: 1px solid #ddd;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.like-section button {
  background: #007bff;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.like-section button:hover {
  background: #0056b3;
}

.like-section button.liked {
  background: #dc3545;
}

.like-section button.liked:hover {
  background: #c82333;
}

.status {
  font-size: 12px;
}

.connected {
  color: #28a745;
}

.disconnected {
  color: #dc3545;
}
</style> 