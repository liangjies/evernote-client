package router

import (
	"evernote-client/api/v1"
	"evernote-client/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("logout", v1.Logout) //退出登录
	}
}
