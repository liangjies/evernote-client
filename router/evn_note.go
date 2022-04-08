package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitNoteRouter(Router *gin.RouterGroup) {
	NoteRouter := Router.Group("notes")
	{
		NoteRouter.GET("list/:id", v1.GetNotes)    // 获取笔记本笔记列表
		NoteRouter.GET("list/all", v1.GetAllNotes) //获取所有笔记
		NoteRouter.POST("add", v1.CreateNote)      //用户新建笔记
		NoteRouter.POST("update", v1.UpdateNote)   //用户修改笔记
		NoteRouter.DELETE(":id", v1.DeleteNote)    //用户删除笔记
		NoteRouter.GET(":id", v1.GetNoteById)      //用户根据id获取笔记详情
		NoteRouter.POST("search", v1.SearchNote)   //用户搜索笔记
	}
}
