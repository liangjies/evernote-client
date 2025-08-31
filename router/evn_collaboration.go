package router

import (
	v1 "evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCollaborationRouter(Router *gin.RouterGroup) {
	CollaborationRouter := Router.Group("collaboration")
	{
		// WebSocket协同编辑连接
		CollaborationRouter.GET("/ws", v1.CollaborationWebSocket)

		// 获取协同编辑状态
		CollaborationRouter.GET("/status/:noteId", v1.GetCollaborationStatus)

		// 获取笔记最终内容
		CollaborationRouter.GET("/content/:noteId", v1.GetNoteFinalContent)
	}
}
