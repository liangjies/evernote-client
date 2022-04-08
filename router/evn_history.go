package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitHistoryRouter(Router *gin.RouterGroup) {
	HistoryRouter := Router.Group("history")
	{
		HistoryRouter.GET(":id", v1.GetHistories)        //获取笔记历史版本
		HistoryRouter.POST("recover", v1.RecoverHistory) //恢复笔记历史版本
	}
}
