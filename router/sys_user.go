package router

import (
	"evernote-client/api/v1"
	"evernote-client/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("").Use(middleware.OperationRecord())
	{
		// UserRouter.POST("register", v1.Register)                 // 用户注册账号
		// UserRouter.POST("changePassword", v1.ChangePassword)     // 用户修改密码
		// UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		// UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		// UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		// UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
		//UserRouter.GET("Login", v1.Login) // 用户修改密码
		UserRouter.GET("notebooks", v1.GetNotebooks)
		UserRouter.POST("notebooks", v1.CreateNotebook)
		UserRouter.PATCH("notebooks/:id", v1.UpdateNotebook)
		UserRouter.DELETE("notebooks/:id", v1.DeleteNotebook)

		UserRouter.GET("notes/from/:id", v1.GetNotes)
		UserRouter.POST("notes/to/:id", v1.CreateNote)
		UserRouter.PATCH("notes/:id", v1.UpdateNote)
		UserRouter.DELETE("notes/:id", v1.DeleteNote)

		UserRouter.GET("notes/trash", v1.GetTrashs)
		UserRouter.DELETE("/notes/confirm/:id", v1.DeleteTrash)
		UserRouter.PATCH("/notes/revert/:id", v1.RevertNote)

	}
}
