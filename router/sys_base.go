package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		//BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.GET("auth", v1.GetAuth)
	}
	return BaseRouter
}
