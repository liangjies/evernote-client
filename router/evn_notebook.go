package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitNoteBookRouter(Router *gin.RouterGroup) {
	NoteBookRouter := Router.Group("notebook")
	{
		NoteBookRouter.GET("get", v1.GetNotebooks)          // 获取笔记本
		NoteBookRouter.POST("add", v1.CreateNotebook)       //新建笔记本
		NoteBookRouter.PATCH(":id", v1.UpdateNotebook)      //修改笔记本
		NoteBookRouter.DELETE("del/:id", v1.DeleteNotebook) //删除笔记本
	}
}
