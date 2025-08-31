package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/response"
	"evernote-client/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var crdtService = service.NewCRDTService()

// 协同编辑WebSocket连接
func CollaborationWebSocket(c *gin.Context) {
	//来源
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            // 基础来源校验：仅允许同源与常见本地开发域名，生产可改为配置白名单
            origin := r.Header.Get("Origin")
            if origin == "" {
                return true
            }
            // 允许同源、常见本地地址
            allowed := []string{
                "http://localhost:5173",
                "http://127.0.0.1:5173",
                "http://localhost:8080",
                "http://127.0.0.1:8080",
                "http://localhost:8888",
                "http://127.0.0.1:8888",
            }
            for _, a := range allowed {
                if origin == a {
                    return true
                }
            }
            return false
        },
    }

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.LOG.Error("WebSocket升级失败: " + err.Error())
		return
	}
	defer conn.Close()

	noteIDStr := c.Query("noteId")
	userIDStr := c.Query("userId")
	username := c.Query("username")

	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		global.LOG.Error("笔记ID解析失败: " + err.Error())
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		global.LOG.Error("用户ID解析失败: " + err.Error())
		return
	}

	crdtService.JoinCollaboration(uint(noteID), uint(userID), username, conn)

	operations, err := crdtService.GetNoteOperations(uint(noteID))
	if err == nil {
		conn.WriteJSON(map[string]interface{}{
			"type":       "history",
			"operations": operations,
		})
	}

	activeUsers, err := crdtService.GetActiveUsers(uint(noteID))
	if err == nil {
		conn.WriteJSON(map[string]interface{}{
			"type":        "active_users",
			"activeUsers": activeUsers,
		})
	}

	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			global.LOG.Error("读取WebSocket消息失败: " + err.Error())
			break
		}

		messageType, ok := message["type"].(string)
		if !ok {
			continue
		}

		switch messageType {
		case "operation":
			// 处理CRDT操作
			if opType, ok := message["opType"].(string); ok {
				if position, ok := message["position"].(float64); ok {
					if content, ok := message["content"].(string); ok {
						crdtService.HandleOperation(uint(noteID), uint(userID), opType, int(position), content)
					}
				}
			}
		case "ping":
			// 心跳检测
			conn.WriteJSON(map[string]interface{}{
				"type": "pong",
			})
		case "leave":
			// 用户主动离开
			if sessionID, ok := message["sessionId"].(string); ok {
				crdtService.LeaveCollaboration(uint(noteID), sessionID)
			}
			return
		}
	}

	crdtService.LeaveCollaboration(uint(noteID), "")
}

// 获取笔记的协同编辑状态
func GetCollaborationStatus(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	activeUsers, err := crdtService.GetActiveUsers(uint(noteID))
	if err != nil {
		response.FailWithMessage("获取协同编辑状态失败", c)
		return
	}

	operations, err := crdtService.GetNoteOperations(uint(noteID))
	if err != nil {
		response.FailWithMessage("获取操作历史失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"activeUsers": activeUsers,
		"operations":  operations,
	}, c)
}

// 获取笔记的最终内容（应用所有操作后）
func GetNoteFinalContent(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	var note model.EvnNote
	err = global.DB.Where("id = ?", noteID).First(&note).Error
	if err != nil {
		response.FailWithMessage("笔记不存在", c)
		return
	}

	operations, err := crdtService.GetNoteOperations(uint(noteID))
	if err != nil {
		response.FailWithMessage("获取操作历史失败", c)
		return
	}

	finalContent := crdtService.ApplyOperationsToContent(note.Content, operations)

	response.OkWithData(map[string]interface{}{
		"noteId":          noteID,
		"originalContent": note.Content,
		"finalContent":    finalContent,
		"operations":      operations,
	}, c)
}
