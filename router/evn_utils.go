package router

import (
	"evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUtilsRouter(Router *gin.RouterGroup) {
	UtilsRouter := Router.Group("")
	{
		UtilsRouter.POST("/upload", v1.UploadFile) // 上传文件
	}
}
