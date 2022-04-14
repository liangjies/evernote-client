package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTrashRouter(Router *gin.RouterGroup) {
	TrashRouter := Router.Group("trash")
	{
		TrashRouter.GET("all", v1.GetTrashs)              //获取所有废纸篓笔记
		TrashRouter.GET(":id", v1.GetTrashById)           //用户根据id获取废纸篓笔记详情
		TrashRouter.DELETE("confirm/:id", v1.DeleteTrash) //用户彻底删除笔记
		TrashRouter.PATCH("revert/:id", v1.RevertNote)    //用户从废纸篓恢复笔记
	}
}
