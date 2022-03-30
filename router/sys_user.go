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
		UserRouter.POST("user/logout", v1.Logout)

		UserRouter.GET("notebook/get", v1.GetNotebooks)
		UserRouter.POST("notebook/add", v1.CreateNotebook)
		UserRouter.PATCH("notebooks/:id", v1.UpdateNotebook)
		UserRouter.DELETE("notebook/del/:id", v1.DeleteNotebook)

		UserRouter.GET("notes/list/:id", v1.GetNotes)
		UserRouter.GET("/notes/list/all", v1.GetAllNotes)
		UserRouter.POST("notes/add", v1.CreateNote)
		UserRouter.POST("notes/update", v1.UpdateNote)
		UserRouter.DELETE("notes/:id", v1.DeleteNote)
		UserRouter.GET("notes/:id", v1.GetNoteById)

		UserRouter.GET("trash/all", v1.GetTrashs)
		UserRouter.GET("trash/:id", v1.GetTrashById)

		UserRouter.DELETE("/trash/confirm/:id", v1.DeleteTrash)
		UserRouter.PATCH("/trash/revert/:id", v1.RevertNote)

	}
}
